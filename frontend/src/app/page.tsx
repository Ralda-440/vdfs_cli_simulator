'use client';

import { useEffect, useRef, useState } from 'react';

import HeaderCli from '@/components/header-cli';
import { Box, Text, Textarea } from '@chakra-ui/react';
import { requestRest } from '@/services/request.service';

function Terminal () {
  const [inputValue, setInputValue] = useState<string>('');
  const [outputValue, setOutputValue] = useState<string>('');
  const textAreaRef = useRef<HTMLTextAreaElement>(null);
  const outputRef = useRef<HTMLDivElement>(null);
  const [comandos, setComandos] = useState<string[]>([]);
  const [placeHolderInput, setPlaceHolderInput] = useState<string>('Ingrese Comando(s)');
  const [isDisabledInput, setIsDisabledInput] = useState<boolean>(false);

  //UseEffect cuando cambiar el valor de comandos
  useEffect(() => {
    // Si no hay comandos, no hacer nada
    if (!comandos.length) {
      setPlaceHolderInput('Ingrese Comando(s)');
      setIsDisabledInput(false);
      return;
    }
    setPlaceHolderInput('Ejecutando Comando(s)...');
    setIsDisabledInput(true);
    // Obtener el primer comando de la lista de comandos y eliminarlo
    const comando = comandos.shift();
    // Fetch para enviar el comando al servidor
    requestRest('POST', '/inputCLI', {input: comando})
      .then((data) => {
        // Agregar los errores al output
        //console.log(data);
        setOutputValue((prevOutputValue) => {
          const newOutputValue = `${prevOutputValue}${
            data.errs.map((element : Error_)  => {
              return element.msg;
            }).join('\n')
          }`+`${data.output != "" ? '\n'+data.output : ''}`+'\n';
          //Guardar outputValue en el LocalStorage
          localStorage.setItem('outputValue', newOutputValue);
          //Si hay errores, no agregar los comandos restantes
          if (data.errs.length > 0) {
            setComandos([]);
            return newOutputValue;
          }
          // Agregar los comandos restantes al estado
          setComandos([...comandos]);
          return newOutputValue;
        });
      })
      .catch((error) => {
        setOutputValue((prevOutputValue) => {
          // Agregar el valor de error al output
          const newOutputValue = `${prevOutputValue}>${error}\n`;
          //Guardar outputValue en el LocalStorage
          localStorage.setItem('outputValue', newOutputValue);
          return newOutputValue;
        });
      });
  }, [comandos]);

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

  const addOuputValue = () => {
    setOutputValue((prevOutputValue) => {
      const newOutputValue = `${prevOutputValue}>${inputValue}\n`;
      //Guardar outputValue en el LocalStorage
      localStorage.setItem('outputValue', newOutputValue);
      return newOutputValue;
    });
  };

  const handleOnInput = (event: React.FormEvent<HTMLTextAreaElement>) => {
    event.currentTarget.style.height = 'auto';
    event.currentTarget.style.height = event.currentTarget.scrollHeight + 'px';
  };

  type Error_ = {
    msg: string;
    linea: number;
    columna: number;
  }

  const handleKeyDown = (event: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (event.key === 'Enter' && !event.shiftKey) {
      // Realizar la función para Enter sin Shift
      event.preventDefault(); // Evita el comportamiento predeterminado del Enter
      // Agregar el valor de input a output
      addOuputValue();
      // Copia de input
      const inputValueCopy = inputValue;
      // Limpiar el valor de input
      setInputValue('');
      event.currentTarget.value = '';
      handleOnInput(event);
      // Si el input está vacío, no enviar nada
      if (!inputValueCopy.trim()) {
        return;
      }
      //Realizar el split de los comandos con salto de linea
      const comandos_ = inputValueCopy.split('\n');
      //Agregar los comandos al estado
      setComandos(comandos_);
    }
  };

  useEffect(() => {
    textAreaRef.current?.focus();
    if (outputRef.current) {
      outputRef.current.scrollTop = outputRef.current.scrollHeight;
    }
  }, [outputValue,isDisabledInput]); // Scroll to bottom on outputValue change

  return (
    <>
      <HeaderCli setOutputValue={setOutputValue}/>
      <Box
        bg="black"
        color="white"
        height="45em"
        width={'auto'}
        maxWidth={'100%'}
        borderRadius="md"
        p={4}
        overflowY="auto"
        ref={outputRef}
        whiteSpace="pre-wrap"
      >
        <Text
        maxWidth={'50em'}>{outputValue}</Text>
        <Box display="flex" mt={2}>
          <Text color="green">{'>$'}</Text>
          <Textarea
            isDisabled={isDisabledInput}
            placeholder={placeHolderInput}
            ref={textAreaRef}
            value={inputValue}
            bg={'transparent'}
            style={{
              caretColor: 'white',
            }}
            onKeyDown={handleKeyDown}
            onChange={handleInputChange}
            width={'100%'}
            height={'auto'}
            padding={0}
            paddingLeft={2}
            variant="unstyled"
            resize="none"
            onInput={handleOnInput}
          ></Textarea>
        </Box>
      </Box>
    </>
  );
};

export default Terminal;
