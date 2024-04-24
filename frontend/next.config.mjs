export default {
  async rewrites() {
    return [
      {
        source: '/api/:path*',
        destination: 'http://localhost:8080/api/:path*', // Adjust to your Go server address
      },
    ];
  },
  images: {
    domains: ['upload.wikimedia.org'], // Add your image domains here
  },
};