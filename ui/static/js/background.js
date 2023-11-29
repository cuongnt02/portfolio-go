var canvas = /** @type{HTMLCanvasElement} */(document.getElementById('background-canvas'));

var ctx = canvas.getContext('2d');

canvas.width = window.innerWidth;
canvas.height = window.innerHeight;


var config = {
    particleNumber: 50,
    maxParticleSize: 10,
    maxSpeed: 20,
    colorVariation: 50
};

var colorPallete = {
    bg: {r:0,g:0,b:0},
    matter: [
        {r:255,g:255,b:255}
    ]
};

var particles = [],
    centerX = canvas.width / 2,
    centerY = canvas.height / 2,
    drawBg;

drawBg = (ctx, color) => {
    ctx.fillStyle = "rgb(" + color.r + ","
        + color.g + "," + color.b + ")";
    ctx.fillRect(0, 0, canvas.width, canvas.height);
};

var Particle = function(x, y) {
    this.x = x || Math.round(Math.random() * canvas.width);
    this.y = y || Math.round(Math.random() * canvas.height);
    this.r = config.maxParticleSize;
    this.c = "rgb(" + colorPallete.matter[0].r + ","
        + colorPallete.matter[0].g + "," + colorPallete.matter[0].b + ")";
    this.s = Math.pow(Math.ceil(Math.random() * config.maxSpeed), .7);
    this.d = Math.random() > Math.random() / 2 ? 90 : 0;
};

var drawParticle = (x, y, r, c) => {
    ctx.beginPath(); 
    ctx.fillStyle = c;
    ctx.arc(x, y, r, 0, 2*Math.PI, false);
    ctx.fill();
    ctx.closePath();
};

var cleanUpArray = function() {
    particles = particles.filter((p) => {
        return (p.x > -100 && p.y > -100);
    });
}

var updateParticleModel = function(p) {
    p.d == 0 ? p.y += p.s : p.y -= p.s;
    return p;
}

var initParticles = function(numParticles) {
    for (let i = 0; i < numParticles; i++) {
        particles.push(new Particle(Math.round(Math.random()*canvas.width), 0));
        particles.forEach((p) => {
            drawParticle(p.x, p.y, p.r, p.c)
        });
    }
}

window.requestAnimFrame = (function() {
    return window.requestAnimationFrame ||
        function(callback) {
            window.setTimeout(callback, 1000 / 60);
        };
})();

var frame = function() {
    drawBg(ctx, colorPallete.bg);
    particles.map((p) => {
        return updateParticleModel(p);
    });
    particles.forEach((p) => {
        drawParticle(p.x, p.y, p.r, p.c);
    });
    window.requestAnimFrame(frame);
}

document.body.addEventListener("wheel", function(event) {
    if (window.scrollY === 0 || window.scrollY >= window.scrollMaxY){
        return;
    }
    cleanUpArray();
    initParticles(config.particleNumber);
})

frame();

initParticles(config.particleNumber);


