// const path = require('path');
import { path } from 'path';
config = {
    entry: './index.js',
    output: {
        path: path.resolve(_dirname, 'dist'),
        filename: 'index.js'
    }
}
module.exports = config;