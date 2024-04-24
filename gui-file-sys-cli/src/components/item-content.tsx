'use client'

import { Box,Text ,useDisclosure,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalBody,
  ModalCloseButton,
} from '@chakra-ui/react';
import { BsDeviceHddFill } from "react-icons/bs";
import { AiFillFileText } from "react-icons/ai";
import { RiDeviceFill } from "react-icons/ri";
import { GoFileDirectoryFill } from "react-icons/go";
import { useState, useEffect ,useRef} from 'react';
import Login  from '@/components/login';
 
const ItemContent = ({nombre,tipo,setContent,fetchExplorer,setIsDisabledLogout}:ItemContentProps) => {
    const [selected, setSelected] = useState<boolean>(false);
    const [hovered, setHovered] = useState<boolean>(false);
    const ref = useRef<HTMLDivElement>(null);
    const { isOpen, onOpen, onClose } = useDisclosure()

    const handleMouseEnter = () => {
        setHovered(true);
    };

    const handleMouseLeave = () => {
        setHovered(false);
    };

    const handleDoubleClick = async () => { 
        if (tipo === "partition") {
            //router.push('/explorer/login');
            const response = await fetch('http://localhost:4005/loginActivo', {
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
            return;
        } else if (tipo === "rep ") {
            return;
        }
        const path = localStorage.getItem('path');
        localStorage.setItem('path', `${path}${path === "/" ? "" : "/"}${nombre}`);
        fetchExplorer();
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
        width={'10%'} 
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
            tipo === "dir" ? <GoFileDirectoryFill size={100} /> : null
        }
        <Text
            userSelect={'none'}
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
   </>
  );
}

type ItemContentProps = {
    nombre: string;
    tipo: string;
    setContent: any;
    fetchExplorer: any;
    setIsDisabledLogout: any;
}

export default ItemContent;