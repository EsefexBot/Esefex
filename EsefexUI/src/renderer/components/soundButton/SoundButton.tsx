import {
  ActionIcon,
  Button,
  Flex,
  Grid,
  Group,
  Image,
  Text,
} from '@mantine/core';
import { FC } from 'react';
import { Sound } from '../../models/Sound';
import './SoundButton.css';
import { IconPlayerPlay } from '@tabler/icons-react';
import axios from 'axios';
import config from '../../config.json';
import {
  showErrorNotification,
  showSuccessNotification,
} from '../notifications';

interface SoundButtonProps {
  sound: Sound;
  serverId: string;
}

export const SoundButton: FC<SoundButtonProps> = ({ sound, serverId }) => {
  const playSound = () => {
    axios
      .post(
        `${config.apiUrl}api/playsound/247763762298355712/${serverId}/${sound.id}`,
      )
      .catch((error) => {
        showErrorNotification(error);
      });
  };

  return (
    <div className="soundButton">
      <Flex h={'100%'} className="soundButton__grid" align={'center'}>
        <Image src={sound.icon} className="soundButton__image" />
        <Text truncate="end" className="soundButton__text">
          {sound.name}
        </Text>
        <ActionIcon variant="subtle" color="#fff" onClick={() => playSound()}>
          <IconPlayerPlay />
        </ActionIcon>
      </Flex>
    </div>
  );
};
