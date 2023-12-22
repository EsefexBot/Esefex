import { Button } from "@mantine/core";
import { FC } from "react";
import { Sound } from "../../models/Sound";

interface SoundButtonProps {
    sound: Sound;
}

export const SoundButton: FC<SoundButtonProps> = ({sound}) => {
    return (
        <>
            <Button w={100} h={100} bg="#1A1A1A" radius={25}><img src={sound.icon}/></Button>
        </>
    );
}