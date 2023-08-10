import { useState, useEffect } from 'react';
import axios from '../utils/axios';
import Nav from '../components/Nav';
import { useRouter } from 'next/router'; // Importing this to use for redirection
import { AxiosError } from 'axios';

interface User {
  username: string;
  id: string; // If the id is a number, you can replace 'string' with 'number'
}

const Profile: React.FC = () => {
  const [user, setUser] = useState<User | null>(null);
  const router = useRouter();

  useEffect(() => {
    const fetchProfile = async () => {
      try {
        // Adding withCredentials to ensure cookies are sent with the request
        const response = await axios.get('/api/v1/profile', { withCredentials: true });
        setUser(response.data.user);
      } catch (error) {
        const axiosError = error as AxiosError; // Type assertion
        console.error('Profile Lookup Error:', axiosError?.response?.data);
    }
    };

    fetchProfile();
  }, [router]); // Added 'router' to the dependency array as it's used inside the useEffect

  return (
    <>
      <Nav />
      <div>
        {user ? (
          <div>
            <p>Username: {user.username}</p>
            <p>ID: {user.id}</p>
          </div>
        ) : (
          <p>Loading...</p>
        )}
      </div>
    </>
  );
}

export default Profile;
