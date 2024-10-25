import ChildButton from "./components/ChildButton";
import { useTextContext } from "./context/TextContext";

const App = () => {
  const { text, textareaRef } = useTextContext();

  return (
    <div>
      <textarea ref={textareaRef} value={text} readOnly rows={10} cols={30} />
      <div>
        <ChildButton lineText="Línea desde Componente 1" />
        <br />
        <ChildButton lineText="Línea desde Componente 2" />
      </div>
    </div>
  );
};

export default App;
