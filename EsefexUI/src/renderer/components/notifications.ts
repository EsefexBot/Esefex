import { showNotification } from "@mantine/notifications";

const showErrorNotification = () => {
    showNotification({
        title: "Error!",
        message: "Something went wrong!",
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