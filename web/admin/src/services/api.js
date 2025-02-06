import axios from 'axios';

const API_BASE_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api/v1';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add token to requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Handle response errors
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export const auth = {
  login: (username, password) =>
    api.post('/auth/login', { username, password }),
  refreshToken: () => api.post('/auth/refresh'),
};

export const messages = {
  send: (data) => api.post('/sms/send', data),
  getStatus: (id) => api.get(`/sms/status/${id}`),
  getHistory: (params) => api.get('/sms/history', { params }),
};

export const operators = {
  list: () => api.get('/operators'),
  add: (data) => api.post('/operators', data),
  update: (id, data) => api.put(`/operators/${id}`, data),
  delete: (id) => api.delete(`/operators/${id}`),
};

export const monitoring = {
  getMetrics: () => api.get('/monitoring/metrics'),
  getStatus: () => api.get('/monitoring/status'),
};

export const settings = {
  get: () => api.get('/settings'),
  update: (data) => api.put('/settings', data),
};

export default api; 