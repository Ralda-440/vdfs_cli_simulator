'use client';

import { useState, useRef, useEffect } from 'react';
import { Box, Text,Textarea,ChakraProvider} from '@chakra-ui/react';

const Terminal: React.FC = () => {
  const [inputValue, setInputValue] = useState<string>('');
  const [outputValue, setOutputValue] = useState<string>('');
  const textAreaRef = useRef<HTMLTextAreaElement>(null);
  const outputRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    // Cargar outputValue del LocalStorage
    const outputValue = localStorage.getItem('outputValue');
    if (outputValue) {
      setOutputValue(outputValue);
    }
  }, []);

  const handleInputChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setInputValue(e.target.value);
  };

  const handleSubmit = () => {
    const newOutputValue = `${outputValue}>${inputValue}\n`;
    setOutputValue(newOutputValue);
    //Guardar outputValue en el LocalStorage
    localStorage.setItem('outputValue', newOutputValue);
    setInputValue('');
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (event.key === "Enter" && !event.shiftKey) {
      // Realizar la función para Enter sin Shift
      event.preventDefault(); // Evita el comportamiento predeterminado del Enter
      // Enviar el comando
      handleSubmit();
      event.currentTarget.value = '';
      event.currentTarget.style.height = 'auto';
      event.currentTarget.style.height = event.currentTarget.scrollHeight + 'px';
    }
  };

  useEffect(() => {
    textAreaRef.current?.focus();
    if (outputRef.current) {
      outputRef.current.scrollTop = outputRef.current.scrollHeight;
    }
  }, [outputValue]); // Scroll to bottom on outputValue change

  return (
    <Box
      bg="black"
      color="white"
      height="45em"
      width={'100%'}
      borderRadius="md"
      p={4}
      overflowY="auto"
      ref={outputRef}
      whiteSpace="pre-wrap"
    >
      <Text
        maxWidth={'90em'}
      >
        {outputValue}
      </Text>
      <Box display="flex" mt={2}>
        <Text color="green"
        >{'>$'}</Text>
        <ChakraProvider>
        <Textarea 
          ref={textAreaRef}
          placeholder='Ingrese Comando(s)'
          bg={'transparent'}
          style={{
            caretColor: 'white',
          }}
          onKeyDown={handleKeyDown}
          onChange={handleInputChange}
          width={'100%'}
          value={inputValue}
          padding={0}
          paddingLeft={2}
          variant='unstyled'
          resize='none'
          onInput={(e) => {
            e.currentTarget.style.height = 'auto';
            e.currentTarget.style.height = e.currentTarget.scrollHeight + 'px';
          }}
          ></Textarea>
        </ChakraProvider>
      </Box>
    </Box>
  );
};

export default Terminal;
