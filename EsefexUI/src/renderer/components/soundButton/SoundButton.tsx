import { Button } from '@mantine/core';
import { Sound } from '../../models';

interface SoundButtonProps {
  sound: Sound;
}

function SoundButton(props: SoundButtonProps) {
  const { sound } = props;

  return (
    <Button w={100} h={100}>
      <img src={sound.icon} alt="test" />
    </Button>
  );
}

export default SoundButton;
