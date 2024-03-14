const http = require('http');
const fs = require('fs');

const PORT = process.env.PORT || 3000;
const FILE_PATH = 'quotes.txt';

const server = http.createServer((req, res) => {
    fs.readFile(FILE_PATH, 'utf8', (err, data) => {
        if (err) {
            res.writeHead(500);
            return res.end('Error reading file');
        }

        const strings = data.split('\n').filter(Boolean);
        const randomString = strings[Math.floor(Math.random() * strings.length)];

        res.writeHead(200, { 'Content-Type': 'text/plain' });
        res.end(randomString);
    });
});

server.listen(PORT, () => {
    console.log(`Server running on port ${PORT}`);
});