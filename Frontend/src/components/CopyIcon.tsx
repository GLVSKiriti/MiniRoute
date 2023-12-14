import copyIcon from "../assets/copy.png";
import { message } from "antd";

type copyIconProps = {
  copyText: string;
  visibility: "visible" | "hidden";
};

/**
 * Imp Note: Here CopyIcon position is absolute,
 * So its parents position should be non static(use relative)
 */
function CopyIcon(props: copyIconProps) {
  const handleCopy = async (text: string) => {
    try {
      await navigator.clipboard.writeText(text);
      message.success("Text copied to clipboard");
    } catch (err) {
      message.error("Failed to copy text");
    }
  };

  return (
    <span
      style={{
        position: "absolute",
        right: "0px",
        height: "20px",
        width: "20px",
        borderRadius: "50%",
        background: "white",
        padding: "5px",
        boxShadow: "0 15px 16.83px 0.17px rgba(0, 0, 0, 0.05)",
        cursor: "pointer",
        visibility: props.visibility,
      }}
      onClick={() => handleCopy(props.copyText)}
    >
      <img src={copyIcon} height="100%" width="100%" />
    </span>
  );
}

export default CopyIcon;
