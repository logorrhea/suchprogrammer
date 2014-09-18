var branches = [];

// a = angle, l = length, d = depth
var seed = {i: 0, x: 420, y:600, a: 0, l: 130, d:0};

var da = 0.5;       // Angle delta
var dl = 0.8;       // Length delta (factor)
var ar = 0.7;       // Randomness
var maxDepth = 10;

var branch = function(b) {
    var end = endPt(b), daR, newB;
    branches.push(b);

    if (b.d === maxDepth)
        return;

    var branchCount = numBranches();
    for (var i = 0; i < branchCount; i++) {
        daR = ar * Math.random() - ar * 1/branchCount*2;
    }

    // Left branch
    daR = ar * Math.random() - ar * 0.5;
    newB = {
        i: branches.length,
        x: end.x,
        y: end.y,
        a: b.a - da + daR,
        l: b.l * dl,
        d: b.d + 1,
        parent: b.i
    };
    branch(newB);

    // Right branch
    daR = ar * Math.random() - ar * 0.5;
    newB = {
        i: branches.length,
        x: end.x,
        y: end.y,
        a: b.a + da + daR,
        l: b.l * dl,
        d: b.d + 1,
        parent: b.i
    };
    branch(newB);
};

var regenerate = function(initialize) {
    branches = [];
    branch(seed);
    initialize ? create() : update();
};

var numBranches = function() {
    return parseInt(Math.random()*10) % 3 + 1;
};

var endPt = function(b) {
    // Return endpoint of branch
    var x = b.x + b.l * Math.sin(b.a);
    var y = b.y - b.l * Math.cos(b.a);
    return {x: x, y: y};
};

var x1 = function(d) { return d.x; };
var y1 = function(d) { return d.y; };
var x2 = function(d) { return endPt(d).x; };
var y2 = function(d) { return endPt(d).y; };

var highlightParents = function(d) {
    var color = d3.event.type === 'mouseover' ? 'green' : '#777';
    var depth = d.d;
    for (var i = 0; i <= depth; i++) {
        d3.select('#id-'+parseInt(d.i)).style('stroke', color);
        d = branches[d.parent];
    }
};

var create = function() {
    console.log('creating repo tree...');
    d3.select('svg')
        .selectAll('line')
        .data(branches)
        .enter()
        .append('line')
        .attr('x1', x1)
        .attr('y1', y1)
        .attr('x2', x2)
        .attr('y2', y2)
        .style('stroke-width', function(d) { return parseInt(maxDepth + 1 - d.d) + 'px'; })
        .attr('id', function(d) { return 'id-'+d.i; })
        .on('mouseover', highlightParents)
        .on('mouseout', highlightParents);
};

var update = function() {
    d3.select('svg')
        .selectAll('line')
        .data(branches)
        .transition()
        .attr('x1', x1)
        .attr('y1', y1)
        .attr('x2', x2)
        .attr('y2', y2);
};

regenerate(true);
