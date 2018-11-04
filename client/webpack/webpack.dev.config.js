var path = require('path');
var webpack = require('webpack');

var config = require('./webpack.base.config.js');

config.mode = 'development';
config.devtool = 'inline-source-map';

config.plugins = config.plugins.concat([
  new webpack.LoaderOptionsPlugin({
    debug: true,
  }),
]);

module.exports = config;
