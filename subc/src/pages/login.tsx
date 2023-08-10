import { useState, ChangeEvent, FormEvent } from 'react';
import axios from '../utils/axios';
import { useUser } from '../context/userContext';
import { useRouter } from 'next/router';
import { AxiosError } from 'axios';

interface FormData {
  username: string;
  password: string;
}



const Login: React.FC = () => {
  const { setUser } = useUser();
  const router = useRouter();
  const [formData, setFormData] = useState<FormData>({
    username: '',
    password: '',
  });

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    });
  };

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await axios.post('/api/v1/login', formData);
      
      // Assuming user is of type any for now, you can replace with a more detailed type if needed
      setUser(response.data.user as any); 
      
      console.log('Login Response:', response.data);
      router.push('/profile');
    } catch (error) {
      const axiosError = error as AxiosError; // Type assertion
      console.error('Login Error:', axiosError?.response?.data);
  }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" name="username" placeholder="Username" onChange={handleChange} />
      <input type="password" name="password" placeholder="Password" onChange={handleChange} />
      <button type="submit">Login</button>
    </form>
  );
}

export default Login;
