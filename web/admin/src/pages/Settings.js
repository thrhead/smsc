import React, { useState, useEffect } from 'react';
import {
  Box,
  Paper,
  Typography,
  TextField,
  Button,
  Grid,
  Switch,
  FormControlLabel,
  Divider,
} from '@mui/material';

export default function Settings() {
  const [settings, setSettings] = useState({
    server: {
      host: '',
      port: 0,
      debug: false,
    },
    smpp: {
      systemId: '',
      password: '',
      port: 0,
      tlsPort: 0,
    },
    monitoring: {
      enabled: true,
      interval: '',
    },
    logging: {
      level: '',
      format: '',
    },
  });

  useEffect(() => {
    // TODO: Fetch settings from API
    setSettings({
      server: {
        host: '0.0.0.0',
        port: 8080,
        debug: false,
      },
      smpp: {
        systemId: 'smsc_gateway',
        password: '********',
        port: 2775,
        tlsPort: 2776,
      },
      monitoring: {
        enabled: true,
        interval: '15s',
      },
      logging: {
        level: 'info',
        format: 'json',
      },
    });
  }, []);

  const handleSave = (event) => {
    event.preventDefault();
    // TODO: Save settings to API
    console.log('Saving settings:', settings);
  };

  const handleChange = (section, field) => (event) => {
    const value = event.target.type === 'checkbox' ? event.target.checked : event.target.value;
    setSettings((prev) => ({
      ...prev,
      [section]: {
        ...prev[section],
        [field]: value,
      },
    }));
  };

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Settings
      </Typography>
      <Paper sx={{ p: 3 }}>
        <form onSubmit={handleSave}>
          <Grid container spacing={3}>
            {/* Server Settings */}
            <Grid item xs={12}>
              <Typography variant="h6" gutterBottom>
                Server Settings
              </Typography>
              <Grid container spacing={2}>
                <Grid item xs={12} md={6}>
                  <TextField
                    fullWidth
                    label="Host"
                    value={settings.server.host}
                    onChange={handleChange('server', 'host')}
                  />
                </Grid>
                <Grid item xs={12} md={6}>
                  <TextField
                    fullWidth
                    type="number"
                    label="Port"
                    value={settings.server.port}
                    onChange={handleChange('server', 'port')}
                  />
                </Grid>
                <Grid item xs={12}>
                  <FormControlLabel
                    control={
                      <Switch
                        checked={settings.server.debug}
                        onChange={handleChange('server', 'debug')}
                      />
                    }
                    label="Debug Mode"
                  />
                </Grid>
              </Grid>
            </Grid>

            <Grid item xs={12}>
              <Divider />
            </Grid>

            {/* SMPP Settings */}
            <Grid item xs={12}>
              <Typography variant="h6" gutterBottom>
                SMPP Settings
              </Typography>
              <Grid container spacing={2}>
                <Grid item xs={12} md={6}>
                  <TextField
                    fullWidth
                    label="System ID"
                    value={settings.smpp.systemId}
                    onChange={handleChange('smpp', 'systemId')}
                  />
                </Grid>
                <Grid item xs={12} md={6}>
                  <TextField
                    fullWidth
                    type="password"
                    label="Password"
                    value={settings.smpp.password}
                    onChange={handleChange('smpp', 'password')}
                  />
                </Grid>
                <Grid item xs={12} md={6}>
                  <TextField
                    fullWidth
                    type="number"
                    label="SMPP Port"
                    value={settings.smpp.port}
                    onChange={handleChange('smpp', 'port')}
                  />
                </Grid>
                <Grid item xs={12} md={6}>
                  <TextField
                    fullWidth
                    type="number"
                    label="SMPP TLS Port"
                    value={settings.smpp.tlsPort}
                    onChange={handleChange('smpp', 'tlsPort')}
                  />
                </Grid>
              </Grid>
            </Grid>

            <Grid item xs={12}>
              <Divider />
            </Grid>

            {/* Monitoring Settings */}
            <Grid item xs={12}>
              <Typography variant="h6" gutterBottom>
                Monitoring Settings
              </Typography>
              <Grid container spacing={2}>
                <Grid item xs={12}>
                  <FormControlLabel
                    control={
                      <Switch
                        checked={settings.monitoring.enabled}
                        onChange={handleChange('monitoring', 'enabled')}
                      />
                    }
                    label="Enable Monitoring"
                  />
                </Grid>
                <Grid item xs={12} md={6}>
                  <TextField
                    fullWidth
                    label="Collection Interval"
                    value={settings.monitoring.interval}
                    onChange={handleChange('monitoring', 'interval')}
                  />
                </Grid>
              </Grid>
            </Grid>

            <Grid item xs={12}>
              <Divider />
            </Grid>

            {/* Logging Settings */}
            <Grid item xs={12}>
              <Typography variant="h6" gutterBottom>
                Logging Settings
              </Typography>
              <Grid container spacing={2}>
                <Grid item xs={12} md={6}>
                  <TextField
                    fullWidth
                    label="Log Level"
                    value={settings.logging.level}
                    onChange={handleChange('logging', 'level')}
                  />
                </Grid>
                <Grid item xs={12} md={6}>
                  <TextField
                    fullWidth
                    label="Log Format"
                    value={settings.logging.format}
                    onChange={handleChange('logging', 'format')}
                  />
                </Grid>
              </Grid>
            </Grid>

            <Grid item xs={12}>
              <Box sx={{ display: 'flex', justifyContent: 'flex-end', mt: 2 }}>
                <Button type="submit" variant="contained">
                  Save Settings
                </Button>
              </Box>
            </Grid>
          </Grid>
        </form>
      </Paper>
    </Box>
  );
} 