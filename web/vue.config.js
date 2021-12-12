module.exports = {
  outputDir: '../public',
  publicPath: '/public',
  devServer: {
    writeToDisk: true,
  },
  configureWebpack: config => {
    if (process.env.NODE_ENV === 'development') {
      return {
        
      }
    } else {

    }
  }
}