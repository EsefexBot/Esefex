import { ThemeIcon } from '@mantine/core';
import {
  IconCircleFilled,
  IconPlayerPlayFilled,
  IconStarFilled,
} from '@tabler/icons-react';
import axios from 'axios';
import { Sound } from '../../models/Sound';
import './SoundButton.css';
import config from '../../config.json';
import { showErrorNotification } from '../notifications';
import EsefexBadge from '../Badge/EsefexBadge';

interface SoundButtonProps {
  sound: Sound;
  serverId: string;
}

// always name the props "props", helps with collaberation, easier to identify

function SoundButton(props: SoundButtonProps) {
  const { sound, serverId } = props;

  const playSound = () => {
    axios
      .post(
        `${config.apiUrl}api/playsound/247763762298355712/${serverId}/${sound.id}`,
      )
      .catch((error) => {
        showErrorNotification(error);
      });
  };

  // destructure props if needed
  return (
    <div className="sound-button">
      <div className="sound-button-wrapper top-wrapper">
        <img className="sound-icon" src={sound.icon} alt={sound.name} />
        <span className="sound-name">{sound.name}</span>
        <ThemeIcon
          variant="light"
          radius="xs"
          size="lg"
          color="#fff"
          onClick={playSound}
        >
          <IconPlayerPlayFilled style={{ width: '70%', height: '70%' }} />
        </ThemeIcon>
      </div>
      <div className="sound-button-wrapper bottom-wrapper">
        <EsefexBadge type="favourite" onClick={() => null}>
          <ThemeIcon variant="light" radius="lg" size="sm" color="yellow">
            <IconStarFilled style={{ width: '70%', height: '70%' }} />
          </ThemeIcon>
        </EsefexBadge>
        <EsefexBadge type="key-bind" onClick={() => null}>
          <ThemeIcon variant="light" radius="lg" size="sm" color="white">
            <IconCircleFilled style={{ width: '70%', height: '70%' }} />
          </ThemeIcon>
          <span>CTRL + X</span>
        </EsefexBadge>
      </div>
    </div>
  );
}

export default SoundButton;
