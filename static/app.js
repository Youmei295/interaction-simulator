// Network Simulator Frontend App
const API_BASE = '/api';

// UI Elements
const tickCounter = document.getElementById('tick-counter');
const nodeCounter = document.getElementById('node-counter');
const edgeCounter = document.getElementById('edge-counter');
const btnTick = document.getElementById('btn-tick');
const btnReset = document.getElementById('btn-reset');
const tooltip = document.getElementById('tooltip');

// Strategy Color Map
const colors = {
    'AlwaysCooperator': '#10b981', // Green
    'AlwaysCheater': '#ef4444',    // Red
    'Copycat': '#3b82f6',          // Blue
    'Grudger': '#f59e0b',          // Amber/Orange
    'Detective': '#8b5cf6'         // Purple
};

// D3 Setup
const container = document.getElementById('graph-container');
let width = container.clientWidth;
let height = container.clientHeight;

const svg = d3.select("#graph-container")
    .append("svg")
    .attr("width", width)
    .attr("height", height)
    .attr("viewBox", [0, 0, width, height]);

// Physics Simulation
const simulation = d3.forceSimulation()
    .force("link", d3.forceLink().id(d => d.id).distance(120))
    .force("charge", d3.forceManyBody().strength(-300))
    .force("center", d3.forceCenter(width / 2, height / 2))
    .force("collide", d3.forceCollide().radius(30));

let link = svg.append("g").selectAll(".link");
let node = svg.append("g").selectAll(".node");

// Handle window resize
window.addEventListener('resize', () => {
    width = container.clientWidth;
    height = container.clientHeight;
    svg.attr("width", width).attr("height", height)
       .attr("viewBox", [0, 0, width, height]);
    simulation.force("center", d3.forceCenter(width / 2, height / 2));
    simulation.alpha(0.3).restart();
});

// Fetch and Render State
async function fetchState() {
    try {
        const res = await fetch(`${API_BASE}/state`);
        const data = await res.json();
        updateViz(data);
    } catch (e) {
        console.error("Failed to fetch state", e);
    }
}

// Tick Simulation
async function advanceTick() {
    btnTick.disabled = true;
    try {
        const res = await fetch(`${API_BASE}/tick`, { method: 'POST' });
        const data = await res.json();
        
        // Flash links to show interaction
        svg.selectAll(".link").classed("active", true);
        setTimeout(() => {
            svg.selectAll(".link").classed("active", false);
        }, 300);

        updateViz(data);
    } catch (e) {
        console.error("Failed to advance tick", e);
    } finally {
        btnTick.disabled = false;
    }
}

// Reset Simulation
async function resetSimulation() {
    const topology = document.getElementById('topology-select').value;
    const size = document.getElementById('size-input').value;
    
    try {
        const res = await fetch(`${API_BASE}/reset?topology=${topology}&size=${size}`, { method: 'POST' });
        const data = await res.json();
        updateViz(data);
    } catch (e) {
        console.error("Failed to reset", e);
    }
}

// Update D3 Visualization
function updateViz(data) {
    // Update Stats
    tickCounter.innerText = data.tick;
    nodeCounter.innerText = Object.keys(data.nodes).length;
    edgeCounter.innerText = data.edges.length;

    // Format data for D3
    const nodes = Object.values(data.nodes);
    const links = data.edges.map(e => ({ source: e.source, target: e.target }));

    // Re-bind links
    link = link.data(links, d => `${d.source}-${d.target}`);
    link.exit().remove();
    const linkEnter = link.enter().append("line")
        .attr("class", "link");
    link = linkEnter.merge(link);

    // Re-bind nodes
    node = node.data(nodes, d => d.id);
    node.exit().remove();
    
    const nodeEnter = node.enter().append("g")
        .attr("class", "node")
        .call(drag(simulation));

    nodeEnter.append("circle")
        .attr("r", 14)
        .attr("fill", d => colors[d.strategy])
        .attr("filter", "drop-shadow(0px 0px 8px rgba(255,255,255,0.2))");
        
    nodeEnter.append("text")
        .attr("dy", -20)
        .attr("text-anchor", "middle")
        .attr("fill", "#f8fafc")
        .attr("font-size", "10px")
        .text(d => d.id);

    node = nodeEnter.merge(node);

    // Update existing node properties (like score) seamlessly
    node.select("circle")
        .transition().duration(300)
        .attr("fill", d => colors[d.strategy])
        .attr("r", d => 14 + Math.min(d.score / 5, 10)); // Grow slightly if doing well

    // Tooltip interaction
    node.on("mouseover", (event, d) => {
        tooltip.style.opacity = 1;
        tooltip.innerHTML = `
            <h4>${d.id}</h4>
            <div><strong>Strategy:</strong> ${d.strategy}</div>
            <div><strong>Score:</strong> ${d.score.toFixed(1)}</div>
        `;
    }).on("mousemove", (event) => {
        tooltip.style.left = (event.pageX + 15) + "px";
        tooltip.style.top = (event.pageY - 15) + "px";
    }).on("mouseout", () => {
        tooltip.style.opacity = 0;
    });

    // Update simulation
    simulation.nodes(nodes);
    simulation.force("link").links(links);
    simulation.alpha(0.3).restart();
}

// Tick function for physics
simulation.on("tick", () => {
    link
        .attr("x1", d => d.source.x)
        .attr("y1", d => d.source.y)
        .attr("x2", d => d.target.x)
        .attr("y2", d => d.target.y);

    node
        .attr("transform", d => `translate(${d.x},${d.y})`);
});

// Drag behavior
function drag(simulation) {
    function dragstarted(event) {
        if (!event.active) simulation.alphaTarget(0.3).restart();
        event.subject.fx = event.subject.x;
        event.subject.fy = event.subject.y;
    }
    function dragged(event) {
        event.subject.fx = event.x;
        event.subject.fy = event.y;
    }
    function dragended(event) {
        if (!event.active) simulation.alphaTarget(0);
        event.subject.fx = null;
        event.subject.fy = null;
    }
    return d3.drag()
        .on("start", dragstarted)
        .on("drag", dragged)
        .on("end", dragended);
}

// Bind events
btnTick.addEventListener('click', advanceTick);
btnReset.addEventListener('click', resetSimulation);

// Init
fetchState();
