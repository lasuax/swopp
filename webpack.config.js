var HtmlWebpackPlugin = require('html-webpack-plugin');
var path = require("path");
var webpack = require("webpack");


module.exports = {
    entry: './src/index.js',
    output: {
        path: path.resolve(__dirname, "./dist"),
        filename: 'bundle.js'
    },
    module: {
        rules: [
            {
                test: /\.css$/, 
                use:['style-loader','css-loader','sass-loader']
            },
            {
                test:/\.js$/,
                exclude:/node-modules/,
                use:'babel-loader'
            }
        ]
    },
    externals: {
        //don't bundle the 'react' npm package with our bundle.js
        //but get it from a global 'React' variable
        'react': 'React',
        'react-dom': 'ReactDOM',
        'jquery': 'jQuery'
    },
    devServer: {
        contentBase: path.join(__dirname, "dist"),
        compress: true,
        hot:true
    },
    plugins: [
        new HtmlWebpackPlugin({
            title: 'Swopp',
            // minify: {
            //     collapseWhitespace: true
            // },
            stats:"error-only",
            template: './index.html'
        }),
        new webpack.HotModuleReplacementPlugin(),
        new webpack.NamedModulesPlugin()
    ]
}