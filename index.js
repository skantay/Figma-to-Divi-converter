const express = require('express');
const bodyParser = require('body-parser');

const app = express();
const PORT = 3000;

app.use(bodyParser.json());

app.post('/oauth/callback', (req, res) => {
    
  console.log('Figma OAuth Callback Received:', req.body);
  res.status(200).send('Callback Received');
});

app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});
