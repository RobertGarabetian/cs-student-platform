import API from './api';

interface Credentials {
  email: string;
  password: string;
}

export const login = async (credentials: Credentials) => {
  const response = await API.post('/login', credentials);
  return response.data;
};

export const register = async (data: { name: string; email: string; password: string }) => {
  const response = await API.post('/register', data);
  return response.data;
};
