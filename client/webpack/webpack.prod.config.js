var path = require('path');
var webpack = require('webpack');

var config = require('./webpack.base.config.js');

config.bail = true;
config.mode = 'production';
// config.devtool = '#source-map'

config.plugins = config.plugins.concat([]);

module.exports = config;
