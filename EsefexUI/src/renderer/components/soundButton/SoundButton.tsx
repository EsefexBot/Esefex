import axios from 'axios';
import { Button } from '@mantine/core';
import { Sound } from '../../models/Sound';
import './SoundButton.css';
import config from '../../config.json';
import { showErrorNotification } from '../notifications';

interface SoundButtonProps {
  sound: Sound;
}

// always name the props "props", helps with collaberation, easier to identify

function SoundButton(props: SoundButtonProps) {
  const { sound } = props;

  const playSound = () => {
    axios
      .post(
        `${config.apiUrl}api/playsound/247763762298355712/${sound.serverId}/${sound.id}`,
      )
      .catch((error) => {
        showErrorNotification(error);
      });
  };

  // destructure props if needed
  return (
    <Button w={100} h={100}>
      <img src={sound.icon} alt="test" />
    </Button>
  );
}

export default SoundButton;
