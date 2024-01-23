import { MemoryRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import '@mantine/core/styles.css';
import { Flex, MantineProvider } from '@mantine/core';
import SoundButton from './components/soundButton/SoundButton';
import { Sound } from './models';
import Header from './components/header/Header';
import Theme from './theme';

const TEST_SOUND: Sound = {
  id: '617007869',
  serverId: 'testserver',
  name: 'test',
  icon: 'https://cdn.discordapp.com/emojis/630819109726191617.webp?size=128&quality=lossless',
};

// use loop, even for testing
function Hello() {
  return (
    <>
      <Header />
      <Flex>
        <SoundButton sound={TEST_SOUND} />
        <SoundButton sound={TEST_SOUND} />
        <SoundButton sound={TEST_SOUND} />
        <SoundButton sound={TEST_SOUND} />
      </Flex>
    </>
  );
}

export default function App() {
  return (
    <MantineProvider defaultColorScheme="dark" theme={Theme}>
      <Router>
        <Routes>
          <Route path="/" element={<Hello />} />
        </Routes>
      </Router>
    </MantineProvider>
  );
}
