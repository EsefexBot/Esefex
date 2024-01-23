import { Button } from '@mantine/core';
import { Sound } from '../../models';

interface SoundButtonProps {
  sound: Sound;
}

// always name the props "props", helps with collaberation, easier to identify

function SoundButton(props: SoundButtonProps) {
  // destructure props if needed
  const { sound } = props;
  return (
    <Button w={100} h={100}>
      <img src={sound.icon} alt="test" />
    </Button>
  );
}

export default SoundButton;
