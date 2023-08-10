import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://localhost:8080', // Adjust this if your Go backend is running on a different port
  withCredentials: true // Allow cookies to be sent and received cross-origin
});

export default instance;

export async function fetchUserData() {
  try {
    const response = await axios.get('/api/v1/profile', { withCredentials: true });
    if (response.data.status === 'success') {
      return response.data.user;
    }
    throw new Error(response.data.message);
  } catch (error) {
    console.error('Error fetching user data:', error.message);
    return null;
  }
}