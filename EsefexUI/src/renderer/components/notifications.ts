import { showNotification } from "@mantine/notifications";
import { AxiosError } from "axios";

const showErrorNotification = (error: AxiosError) => {
    showNotification({
        title: "Error!",
        message: error.code == "ERR_NETWORK" ? "Joshua skill issue" : "Something went wrong...",
        autoClose: 5000,
        color: "red",
    });
};

const showSuccessNotification = () => {
    showNotification({
        title: "Success!",
        message: "Everything was saved :)",
        autoClose: 5000,
    });
};

export { showErrorNotification, showSuccessNotification };