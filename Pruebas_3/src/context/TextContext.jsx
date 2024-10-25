import { createContext, useState, useRef, useContext } from "react";

export const TextContext = createContext();

export const TextProvider = ({ children }) => {
  const [text, setText] = useState("");
  const textareaRef = useRef(null);

  const addLine = (line) => {
    setText((prevText) => prevText + line + "\n");
    setTimeout(() => {
      textareaRef.current.scrollTop = textareaRef.current.scrollHeight;
    }, 0);
  };

  return (
    <TextContext.Provider value={{ text, addLine, textareaRef }}>
      {children}
    </TextContext.Provider>
  );
};

export const useTextContext = () => {
  return useContext(TextContext);
};
