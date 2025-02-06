import React, { useState, useEffect } from 'react';
import {
  Box,
  Paper,
  Typography,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Button,
  IconButton,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Snackbar,
  Alert,
} from '@mui/material';
import {
  Edit as EditIcon,
  Delete as DeleteIcon,
  Add as AddIcon,
} from '@mui/icons-material';
import { API_BASE_URL } from '../config';

export default function Operators() {
  const [operators, setOperators] = useState([]);
  const [open, setOpen] = useState(false);
  const [editOperator, setEditOperator] = useState(null);
  const [formData, setFormData] = useState({
    name: '',
    priority: '',
    weight: '',
    maxTps: '',
  });
  const [snackbar, setSnackbar] = useState({
    open: false,
    message: '',
    severity: 'success',
  });

  const fetchOperators = async () => {
    try {
      const response = await fetch(`${API_BASE_URL}/operators/`);
      if (!response.ok) throw new Error('Failed to fetch operators');
      const data = await response.json();
      setOperators(data);
    } catch (error) {
      console.error('Error fetching operators:', error);
      setSnackbar({
        open: true,
        message: 'Failed to load operators',
        severity: 'error',
      });
    }
  };

  useEffect(() => {
    fetchOperators();
  }, []);

  const handleOpen = (operator = null) => {
    setEditOperator(operator);
    setFormData(
      operator
        ? {
            name: operator.name,
            priority: operator.priority,
            weight: operator.weight,
            maxTps: operator.maxTps,
          }
        : {
            name: '',
            priority: '',
            weight: '',
            maxTps: '',
          }
    );
    setOpen(true);
  };

  const handleClose = () => {
    setEditOperator(null);
    setFormData({
      name: '',
      priority: '',
      weight: '',
      maxTps: '',
    });
    setOpen(false);
  };

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSave = async (event) => {
    event.preventDefault();
    try {
      const url = editOperator
        ? `${API_BASE_URL}/operators/${editOperator.id}`
        : `${API_BASE_URL}/operators/`;
      
      const requestBody = {
        name: formData.name,
        priority: parseInt(formData.priority),
        weight: parseInt(formData.weight),
        maxTps: parseInt(formData.maxTps),
      };

      console.log('Sending request to:', url);
      console.log('Request body:', requestBody);

      const response = await fetch(url, {
        method: editOperator ? 'PUT' : 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Accept': 'application/json'
        },
        body: JSON.stringify(requestBody),
      });

      console.log('Response status:', response.status);
      const responseText = await response.text();
      console.log('Response text:', responseText);

      let responseData;
      try {
        responseData = JSON.parse(responseText);
        console.log('Response data:', responseData);
      } catch (e) {
        console.error('Failed to parse response as JSON:', e);
        throw new Error('Invalid response format from server');
      }

      if (!response.ok) {
        throw new Error(responseData.error || `Failed to save operator: ${response.status} ${response.statusText}`);
      }

      setSnackbar({
        open: true,
        message: `Operator ${editOperator ? 'updated' : 'added'} successfully`,
        severity: 'success',
      });
      handleClose();
      fetchOperators();
    } catch (error) {
      console.error('Error saving operator:', error);
      setSnackbar({
        open: true,
        message: error.message || 'Failed to save operator',
        severity: 'error',
      });
    }
  };

  const handleDelete = async (id) => {
    try {
      const response = await fetch(`${API_BASE_URL}/operators/${id}`, {
        method: 'DELETE',
      });

      if (!response.ok) throw new Error('Failed to delete operator');

      setSnackbar({
        open: true,
        message: 'Operator deleted successfully',
        severity: 'success',
      });
      fetchOperators();
    } catch (error) {
      console.error('Error deleting operator:', error);
      setSnackbar({
        open: true,
        message: 'Failed to delete operator',
        severity: 'error',
      });
    }
  };

  const handleSnackbarClose = () => {
    setSnackbar((prev) => ({
      ...prev,
      open: false,
    }));
  };

  return (
    <Box>
      <Box sx={{ display: 'flex', justifyContent: 'space-between', mb: 2 }}>
        <Typography variant="h4">Operators</Typography>
        <Button
          variant="contained"
          startIcon={<AddIcon />}
          onClick={() => handleOpen()}
        >
          Add Operator
        </Button>
      </Box>

      <Paper>
        <TableContainer>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>ID</TableCell>
                <TableCell>Name</TableCell>
                <TableCell>Priority</TableCell>
                <TableCell>Weight</TableCell>
                <TableCell>Max TPS</TableCell>
                <TableCell>Status</TableCell>
                <TableCell>Actions</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {operators.map((operator) => (
                <TableRow key={operator.id}>
                  <TableCell>{operator.id}</TableCell>
                  <TableCell>{operator.name}</TableCell>
                  <TableCell>{operator.priority}</TableCell>
                  <TableCell>{operator.weight}</TableCell>
                  <TableCell>{operator.maxTps}</TableCell>
                  <TableCell>{operator.status}</TableCell>
                  <TableCell>
                    <IconButton
                      color="primary"
                      onClick={() => handleOpen(operator)}
                    >
                      <EditIcon />
                    </IconButton>
                    <IconButton
                      color="error"
                      onClick={() => handleDelete(operator.id)}
                    >
                      <DeleteIcon />
                    </IconButton>
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Paper>

      <Dialog open={open} onClose={handleClose}>
        <form onSubmit={handleSave}>
          <DialogTitle>
            {editOperator ? 'Edit Operator' : 'Add Operator'}
          </DialogTitle>
          <DialogContent>
            <TextField
              autoFocus
              margin="dense"
              name="name"
              label="Name"
              type="text"
              fullWidth
              value={formData.name}
              onChange={handleInputChange}
              required
            />
            <TextField
              margin="dense"
              name="priority"
              label="Priority"
              type="number"
              fullWidth
              value={formData.priority}
              onChange={handleInputChange}
              required
            />
            <TextField
              margin="dense"
              name="weight"
              label="Weight"
              type="number"
              fullWidth
              value={formData.weight}
              onChange={handleInputChange}
              required
            />
            <TextField
              margin="dense"
              name="maxTps"
              label="Max TPS"
              type="number"
              fullWidth
              value={formData.maxTps}
              onChange={handleInputChange}
              required
            />
          </DialogContent>
          <DialogActions>
            <Button onClick={handleClose}>Cancel</Button>
            <Button type="submit" variant="contained">
              Save
            </Button>
          </DialogActions>
        </form>
      </Dialog>

      <Snackbar
        open={snackbar.open}
        autoHideDuration={6000}
        onClose={handleSnackbarClose}
      >
        <Alert
          onClose={handleSnackbarClose}
          severity={snackbar.severity}
          sx={{ width: '100%' }}
        >
          {snackbar.message}
        </Alert>
      </Snackbar>
    </Box>
  );
} 