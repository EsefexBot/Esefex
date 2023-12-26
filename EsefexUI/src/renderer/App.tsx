import { MemoryRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import '@mantine/core/styles.css';
import { MantineProvider,SimpleGrid } from '@mantine/core';
import { theme } from './theme';
import { SoundButton } from './components/soundButton/SoundButton';
import { Sound } from './models/Sound';
import { Header } from './components/header/Header';
import { useEffect, useState } from 'react';
import axios from 'axios';
import config from "./config.json";
import { Notifications } from '@mantine/notifications';
import '@mantine/notifications/styles.css';
import { showErrorNotification } from './components/notifications';

let testSound: Sound = {
  id:"617007869",
  serverId:"testserver",
  name:"Vine Boom Sound Effect",
  icon:"https://cdn.discordapp.com/attachments/248446166927147009/1187897634254508173/tower.png?ex=65988ee2&is=658619e2&hm=b66e6fc79eab2f9958f0129f50cf6e393fbd7796b87058413b455a3842edbbea&"
}

function Hello() {
  const [sounds, setSounds] = useState<Sound[] | null>(null);
  const [serverId, setServerId] = useState<string>("489017101894418444");

  useEffect(() => {
    getData();
  },[]);

  const getData = () => {
    axios.get(config.apiUrl + "api/sounds/" + serverId)
    .then((response) => {
      setSounds(response.data);
    }).catch(() => {
      showErrorNotification();
      setSounds([testSound,testSound])
    });
  };

  const soundButtons = sounds ? sounds.map((sound) =>(
    <SoundButton sound={sound} serverId={serverId} key={sound.id}/>
  )) : (<p>Loading...</p>)

  return (
    <>
      <Header/>
      <SimpleGrid cols={3} spacing={'lg'} m={30} mt={0}>
        {soundButtons}
      </SimpleGrid>
    </>
  );
}

export default function App() {
  return (
    <MantineProvider defaultColorScheme='dark' theme={theme}>
      <Notifications />
      <Router>
        <Routes>
          <Route path="/" element={<Hello />} />
        </Routes>
      </Router>
    </MantineProvider>
  );
}
