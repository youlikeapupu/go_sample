const path = require("path");
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
var htmlPlugin = require("html-webpack-plugin");
var myHtmlPlugin = new htmlPlugin({
    template: "./resources/index.html"
});

module.exports = {
    mode: 'development',
    entry: {
        app: ['@babel/polyfill', './resources/app.js'],
    },
    output: {
        path: path.join( __dirname, "build"),
        filename: "lazy/[name].js",
        publicPath: ""
    },
    module: {
        rules: [{
            test: /\.js$/,
            exclude: /node_modules/,
            use: {
                loader:"babel-loader",
                options:{
                    presets: [
                        "@babel/preset-env",
                        "@babel/preset-react",
                        {"plugins": ["@babel/plugin-proposal-class-properties"]}
                    ],
                }
            },
        }],
    },
    plugins: [
        myHtmlPlugin,
        new  CleanWebpackPlugin({
            cleanOnceBeforeBuildPatterns: ['lazy', 'index.html'],
        }),
    ],
};