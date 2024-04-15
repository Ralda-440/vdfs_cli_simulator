'use client'

import { Box,Text ,useDisclosure,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalBody,
  ModalCloseButton,
} from '@chakra-ui/react';
import { BsDeviceHddFill } from "react-icons/bs";
import { RiDeviceFill } from "react-icons/ri";
import { useState, useEffect ,useRef} from 'react';
import { useRouter } from 'next/navigation';
import Login  from '@/components/login';
 
const ItemContent = ({nombre,tipo,setContent,fetchExplorer}:ItemContentProps) => {
    const [selected, setSelected] = useState<boolean>(false);
    const [hovered, setHovered] = useState<boolean>(false);
    const ref = useRef<HTMLDivElement>(null);
    const router = useRouter();
    const { isOpen, onOpen, onClose } = useDisclosure()

    const handleMouseEnter = () => {
        setHovered(true);
    };

    const handleMouseLeave = () => {
        setHovered(false);
    };

    const handleDoubleClick = () => { 
        if (tipo === "partition") {
            //router.push('/explorer/login');
            onOpen();
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
            tipo === "partition" ? <RiDeviceFill size={100} /> : null
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
            <Login namePart={nombre} />
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
}

export default ItemContent;