import axios from 'axios';
import { Badge, Button } from '@mantine/core';
import { Sound } from '../../models/Sound';
import './SoundButton.css';
import config from '../../config.json';
import { showErrorNotification } from '../notifications';
import EsefexBadge from '../Badge/EsefexBadge';

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
    <div className="sound-button">
      <div className="sound-button-wrapper">
        <img className="sound-icon" src={sound.icon} alt={sound.name} />
        <span className="sound-name">{sound.name}</span>
        <img className="sound-play-button" src={sound.icon} alt={sound.name} />
      </div>
      <div className="sound-button-wrapper">
        <EsefexBadge />
        <EsefexBadge />
      </div>
    </div>
  );
}

export default SoundButton;
