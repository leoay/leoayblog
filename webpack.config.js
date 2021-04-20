const path = require('path');

module.exports = {
    entry: "/resource/js/app.js",
    output: {
        path: path.resolve(__dirname, 'static/js'),
        filename: "app.js"
    },
    module: {
        rules: [
            {
                test: /\.(css|scss)$/, 
                use: [
                        {
                            loader: 'style-loader' 
                        },
                        {
                            loader: 'css-loader'
                        },
                        {
                            loader: 'sass-loader'
                        },
                    ]
            }
        ]
    }
};