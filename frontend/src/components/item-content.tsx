'use client'

import { Box,Text ,useDisclosure,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalBody,
  ModalCloseButton,
  ModalHeader,
} from '@chakra-ui/react';
import { BsDeviceHddFill } from "react-icons/bs";
import { AiFillFileText } from "react-icons/ai";
import { RiDeviceFill } from "react-icons/ri";
import { BiSolidFilePng } from "react-icons/bi";
import { BiSolidFileTxt } from "react-icons/bi";
import { GoFileDirectoryFill } from "react-icons/go";
import { useState, useEffect ,useRef} from 'react';
import Login  from '@/components/login';
 
const ItemContent = ({nombre,tipo,fetchExplorer,setIsDisabledLogout,contentRep}:ItemContentProps) => {
    const [selected, setSelected] = useState<boolean>(false);
    const [hovered, setHovered] = useState<boolean>(false);
    const ref = useRef<HTMLDivElement>(null);
    const { isOpen, onOpen, onClose } = useDisclosure()
    const { isOpen: isOpenContent, onOpen: onOpenContent, onClose: onCloseContent } = useDisclosure()
    const [contentFile, setContentFile] = useState<string>("");

    const handleMouseEnter = () => {
        setHovered(true);
    };

    const handleMouseLeave = () => {
        setHovered(false);
    };

    type Error_ = {
        msg: string;
        linea: number;
        columna: number;
      };

    const handleDoubleClick = async () => { 
        if (tipo === "partition") {
            //router.push('/explorer/login');
            const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/loginActivo`, {
                method: 'GET',
            });
            const data = await response.json();
            if (data.activo) {
                const path = localStorage.getItem('path');
                const disk = path?.split('/')[1];
                const nameDisk = disk?.split('.')[0];
                if (!(data.diskName === nameDisk && data.partName === nombre)) {
                    onOpen();
                    return;
                }
            } else {
                onOpen();
                return;
            }
        } else if (tipo === "file") {
            const path = localStorage.getItem('path');
            const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/contentFile`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    path: path+"/"+nombre,
                }),
            });
            const data = await response.json();
            if ( data.errs !== null && data.errs.length > 0) {
                alert(
                    data.errs.map((err: Error_) => {
                        return err.msg;
                    }).join('\n'),
                );
                return;
            }
            setContentFile(data.content);
            onOpenContent();
            return;
        } else if (tipo === "image") {
            const contentPage = `<html style="height: 100%;"><head><meta name="viewport" content="width=device-width, minimum-scale=0.1"><title>${nombre}</title></head><body style="margin: 0px; height: 100%; background-color: rgb(14, 14, 14);"><img style="display: block;-webkit-user-select: none;margin: auto;background-color: hsl(0, 0%, 90%);transition: background-color 300ms;" src="${contentRep}"></body></html>`;
            const newWindow = window.open();
            newWindow?.document.write(contentPage);
            newWindow?.focus();
            return;
        } else if (tipo === "plain") {
            const contentPage = `<html><head><meta name="color-scheme" content="light dark"></head><body><pre style="word-wrap: break-word; white-space: pre-wrap;">${contentRep}</pre></body></html>`;
            const newWindow = window.open();
            newWindow?.document.write(contentPage);
            newWindow?.focus();
            return;
        }
        const path = localStorage.getItem('path');
        localStorage.setItem('path', `${path}${path === "/" ? "" : "/"}${nombre}`);
        fetchExplorer?.();
    }

    useEffect(() => {
        document.addEventListener('mousedown', (event) => {
            if (ref.current && !ref.current.contains(event.target as Node)) {
                setSelected(false);
            }
        });

        return () => {
            document.removeEventListener('mousedown', (event) => {
                if (ref.current && !ref.current.contains(event.target as Node)) {
                    setSelected(false);
                }
            });
        };
    }, []);
    

  return (
   <>
    <Box
        onDoubleClick={handleDoubleClick}
        ref={ref}
        display={'inline-grid'}
        width={'auto'} 
        justifyItems={'center'}
        onMouseEnter={handleMouseEnter}
        onMouseLeave={handleMouseLeave}
        onClick={() => setSelected(true)}
        bg={selected ? 'gray.400' : hovered ? 'gray.300' : 'transparent'}
    >
        {
            tipo === "disk" ? <BsDeviceHddFill size={100}/> : 
            tipo === "partition" ? <RiDeviceFill size={100} /> :
            tipo === "file" ? <AiFillFileText size={100} /> :
            tipo === "dir" ? <GoFileDirectoryFill size={100} /> :
            tipo === "image" ? <BiSolidFilePng size={100} /> :
            tipo === "plain" ? <BiSolidFileTxt size={100} /> : null
        }
        <Text
            userSelect={'none'}
            maxWidth={'3.8em'}
        >{nombre}</Text>
    </Box>
    <Modal isOpen={isOpen} onClose={onClose}
        size={"full"}
    >
        <ModalOverlay />
        <ModalContent>
          <ModalCloseButton />
          <ModalBody>
            <Login namePart={nombre} onClose={onClose} fetchExplorer={fetchExplorer} setIsDisabledLogout={setIsDisabledLogout} />
          </ModalBody>
        </ModalContent>
      </Modal>
      <Modal isOpen={isOpenContent} onClose={onCloseContent}
        size={"full"}
    >
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>File Content: {nombre}</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Text
                whiteSpace={'pre-wrap'}
            >{contentFile}</Text>
          </ModalBody>
        </ModalContent>
      </Modal>
   </>
  );
}

type ItemContentProps = {
    nombre: string;
    tipo: string;
    setContent?: React.Dispatch<React.SetStateAction<Map<string, string>>>;
    fetchExplorer?: () => Promise<void>;
    setIsDisabledLogout?: React.Dispatch<React.SetStateAction<boolean>>;
    contentRep?: string;
}

export default ItemContent;