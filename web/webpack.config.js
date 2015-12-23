'use strict'
let path = require('path')
let webpack = require('webpack')
let ExtractTextPlugin = require('extract-text-webpack-plugin')

/**
 * @returns {Boolean}
 */
function isProduction() {
  return (
    process.env.NODE_ENV &&
    process.env.NODE_ENV === 'production'
  )
}

/**
 * @returns {Object[]}
 */
function getPlugins() {
  let plugins = [
    new webpack.optimize.DedupePlugin(),
    new webpack.DefinePlugin({
      __DEV__: ! isProduction(),
      'process.env': {
        'NODE_ENV': isProduction() ? "'production'" : "'development'"
      }
    }),
    new ExtractTextPlugin('bundle.css')
  ]

  if (isProduction()) {
    plugins.concat(
      new webpack.optimize.UglifyJsPlugin({
        compress: {
          warnings: false
        }
      })
    )
  } else {
    plugins.push(new webpack.HotModuleReplacementPlugin())
  }

  return plugins
}

/**
 * @returns {String[]}
 */
function getEntry() {
  if (isProduction()) {
    return [
      './src/index'
    ]
  }

  return [
    'webpack-dev-server/client?http://localhost:5000',
    'webpack/hot/only-dev-server',
    './src/index'
  ]
}

module.exports = {
  cache: true,
  context: __dirname,
  entry: getEntry(),
  plugins: getPlugins(),
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js',
    publicPath: '/dist/'
  },
  resolve: {
    moduleDirectories: [
      __dirname + '/src',
      'node_modules'
    ],
    extensions: ['', '.js', '.jsx']
  },
  module: {
    loaders: [
      {
        test: /\.jsx?$/,
        loaders: isProduction() ? ['babel'] : ['react-hot', 'babel'],
        include: [
          path.resolve(__dirname, 'src')
        ],
      }, {
        test: /\.(png|svg|jpe?g|woff)$/,
        loader: 'url-loader?limit=10000'
      }, {
        test: /\.css$/,
        loader: ExtractTextPlugin.extract('style-loader', 'css-loader')
      }
    ]
  }
}
