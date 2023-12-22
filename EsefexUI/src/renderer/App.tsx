import { MemoryRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import '@mantine/core/styles.css';
import { Button, Flex, Grid, MantineProvider,Title } from '@mantine/core';
import { theme } from './theme';
import { SoundButton } from './components/soundButton/SoundButton';
import { Sound } from './models/Sound';
import { Header } from './components/header/Header';

let testSound: Sound = {
  id:"617007869",
  serverId:"testserver",
  name:"test",
  icon:"https://cdn.discordapp.com/emojis/630819109726191617.webp?size=128&quality=lossless"
}

function Hello() {
  return (
    <>
      <Header/>
      <Flex >
        <SoundButton sound={testSound}/>
        <SoundButton sound={testSound}/>
        <SoundButton sound={testSound}/>
        <SoundButton sound={testSound}/>
      </Flex>
      
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
