import { Group, Image, Title } from '@mantine/core';
import './Header.css';

// dont use a react-fragment if only one element is wrapped --> unneccesairy code
function Header() {
  return (
    <Group grow>
      <div>
        <Image
          w={100}
          h={100}
          src="https://cdn.discordapp.com/attachments/248446166927147009/1177018089712402432/esefex.png?ex=658caa04&is=657a3504&hm=236dd01dab7c2356c82afccf9fab8a1bbd823774080c230565cec71fb5f62773&"
        />
      </div>
      <Title className="serverName" mr={20}>
        Server Name
      </Title>
    </Group>
  );
}

export default Header;
