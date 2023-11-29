import { MemoryRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import '@mantine/core/styles.css';
import { Button, MantineProvider,Title } from '@mantine/core';
import { theme } from './theme';

function Hello() {
  return (
    <>
      <Title>ESEFEX</Title>
      <Button color='green'>Test</Button>
    </>
  );
}

export default function App() {
  return (
    <MantineProvider defaultColorScheme='dark' theme={theme}>
      <Router>
        <Routes>
          <Route path="/" element={<Hello />} />
        </Routes>
      </Router>
    </MantineProvider>
  );
}
