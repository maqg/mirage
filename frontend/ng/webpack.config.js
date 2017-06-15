module.exports = {
  entry: './main.ts',

  output: {
    filename: './bundle.js'
  },

  resolve: {
    extensions: ['', '.ts', '.js', '.html']
  },

  module: {
    loaders: [
      {
        test: /\.ts$/,
        loader: 'ts-loader'
      },
      {
        test:/\.css$/,
        loaders:['style','css']
      },
      {
	test: /\.html$/,
	loader: 'raw-loader'
      },
    ]
  }
};
