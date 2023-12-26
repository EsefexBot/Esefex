import { Button, Group, Image, Title } from "@mantine/core";
import { FC } from "react";
import "./Header.css";
import logo from "../../../../assets/ESEFEX.svg";

export const Header: FC = () => {
    return (
        <div className="header">
        <Group grow>
                <div>
                    <Image w={50} h={50} src={logo} fit="contain" />
                </div>
                <Title className="serverName" >Server Name</Title>
            </Group>
        </div>
    );
}