import { Group, Image, Title } from '@mantine/core';
import './Header.css';
import logo from '../../../../assets/ESEFEX.svg';

function Header() {
  return (
    <div className="header">
      <Group grow>
        <div>
          <Image w={50} h={50} src={logo} fit="contain" />
        </div>
        <Title className="serverName">Server Name</Title>
      </Group>
    </div>
  );
}

export default Header;
