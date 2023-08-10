import { useState, ChangeEvent } from 'react';
import axios from '../utils/axios';
import { useRouter } from 'next/router'; 
import { AxiosError } from 'axios';

interface FormData {
  username: string;
  email: string;
  password: string;
}

const Signup: React.FC = () => {
  const [formData, setFormData] = useState<FormData>({
    username: '',
    email: '',
    password: '',
  });
  
  const router = useRouter();

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const response = await axios.post('/api/v1/signup', formData);
      console.log('Signup Response:', response.data);
      router.push('/login'); 
    } catch (error) {
      const axiosError = error as AxiosError; // Type assertion
      console.error('Signup Error:', axiosError?.response?.data);
  }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input type="text" name="username" placeholder="Username" onChange={handleChange} />
      <input type="email" name="email" placeholder="Email" onChange={handleChange} />
      <input type="password" name="password" placeholder="Password" onChange={handleChange} />
      <button type="submit">Signup</button>
    </form>
  );
}

export default Signup;
