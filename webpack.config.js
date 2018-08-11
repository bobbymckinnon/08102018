var path = require('path');
var webpack = require('webpack');

var pathsLookup = [
    path.resolve(__dirname, 'static'),
    path.resolve(__dirname, 'client')
]

module.exports =  {
    entry: ['babel-polyfill', './client/js/app.jsx'],
    output: { path: __dirname, filename: './static/dist/js/bundle.js' },
    mode: 'production',
    module: {
        rules: [
            {
                test: /\.css$/,
                use: [
                    { loader: 'style-loader' },
                    { loader: 'css-loader' }
                ]
            },
            {
                test: /\.jsx?$/,
                exclude: /node_modules/,
                include: pathsLookup,
                loader: 'babel-loader?cacheDirectory'
            },

        ]
    },
    resolve: {
        extensions: ['.js', '.jsx', '.css'],
    }
};
