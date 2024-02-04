const path = require('path');


module.exports = {
  entry: './scripts/updateGasPrice.js',
  target:"node",
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js',
  },
  resolve: {
    mainFields: ["main"]
  }
};