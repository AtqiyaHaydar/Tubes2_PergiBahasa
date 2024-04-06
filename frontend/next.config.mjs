// next.config.mjs
export default {
  async rewrites() {
    return [
      {
        source: '/api/:path*',
        destination: 'http://localhost:8080/api/:path*', // Sesuaikan dengan alamat server Go Anda
      },
    ]
  },
}