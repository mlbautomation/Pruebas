import { useTextContext } from "../context/TextContext";

const ChildButton = ({ lineText }) => {
  const { addLine } = useTextContext();

  return <button onClick={() => addLine(lineText)}>Agregar {lineText}</button>;
};

export default ChildButton;
