'use client';

import { useState } from 'react';

import {
  Alert,
  AlertDescription,
  AlertIcon,
  AlertTitle,
  Avatar,
  Box,
  Button,
  chakra,
  Flex,
  FormControl,
  Heading,
  Input,
  InputGroup,
  InputLeftElement,
  InputRightElement,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalOverlay,
  ModalHeader,
  Stack,
  useDisclosure,
} from '@chakra-ui/react';
import { FaLock, FaUserAlt } from 'react-icons/fa';

const CFaUserAlt = chakra(FaUserAlt);
const CFaLock = chakra(FaLock);

const Login = ({ namePart, onClose, fetchExplorer,setIsDisabledLogout }: LoginProps) => {
  const [showPassword, setShowPassword] = useState(false);
  const [user, setUser] = useState('');
  const [password, setPassword] = useState('');
  const {
    isOpen: isOpenAlert,
    onOpen: onOpenAlert,
    onClose: onCloseAlert,
  } = useDisclosure();
  const [msgAlert, setMsgAlert] = useState('');

  const handleShowClick = () => setShowPassword(!showPassword);

  type Error_ = {
    msg: string;
    linea: number;
    columna: number;
  };

  const handleLogin = async () => {
    const path = localStorage.getItem('path');
    const disk = path?.split('/')[1];
    const nameDisk = disk?.split('.')[0];
    const response = await fetch('http://localhost:4005/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        usr: user,
        pwd: password,
        diskName: nameDisk,
        partName: namePart,
      }),
    });
    const data = await response.json();
    if (data.errs != null && data.errs.length > 0) {
      setMsgAlert(data.errs
        .map((err: Error_) => {
          return err.msg;
        })
        .join('\n'));
      onOpenAlert();
      return;
    }
    localStorage.setItem(
      'path',
      `${path}${path === '/' ? '' : '/'}${namePart}`,
    );
    onClose();
    fetchExplorer?.();
    setIsDisabledLogout?.(false);
  };

  return (
    <>
      <Flex
        flexDirection="column"
        width="50wh"
        height="50vh"
        backgroundColor="transparent"
        justifyContent="center"
        alignItems="center"
      >
        <Stack
          flexDir="column"
          mb="2"
          justifyContent="center"
          alignItems="center"
        >
          <Avatar bg="blackAlpha.900" />
          <Heading color="blackAlpha.900">
            Bienvenido a la Particion {namePart}
          </Heading>
          <Box minW={{ base: '90%', md: '468px' }}>
            <form>
              <Stack
                spacing={4}
                p="1rem"
                backgroundColor="whiteAlpha.900"
                boxShadow="md"
              >
                <FormControl>
                  <InputGroup>
                    <InputLeftElement
                      pointerEvents="none"
                    >
                    <CFaUserAlt color="gray.600" />
                    </InputLeftElement>
                    <Input
                      type="email"
                      placeholder="Usuario"
                      focusBorderColor="blackAlpha.900"
                      value={user}
                      onChange={(e) => setUser(e.target.value)}
                    />
                  </InputGroup>
                </FormControl>
                <FormControl>
                  <InputGroup>
                    <InputLeftElement
                      pointerEvents="none"
                    >
                    <CFaLock color="gray.600" />
                    </InputLeftElement>
                    <Input
                      type={showPassword ? 'text' : 'password'}
                      placeholder="Password"
                      focusBorderColor="blackAlpha.900"
                      value={password}
                      onChange={(e) => setPassword(e.target.value)}
                    />
                    <InputRightElement width="4.5rem">
                      <Button h="1.75rem" size="sm" onClick={handleShowClick}>
                        {showPassword ? 'Hide' : 'Show'}
                      </Button>
                    </InputRightElement>
                  </InputGroup>
                </FormControl>
                <Button
                  borderRadius={0}
                  variant={'outline'}
                  colorScheme="blackAlpha"
                  width="full"
                  color={'blackAlpha'}
                  onClick={handleLogin}
                >
                  Login
                </Button>
              </Stack>
            </form>
          </Box>
        </Stack>
      </Flex>
      <Modal isOpen={isOpenAlert} onClose={onCloseAlert} size={'xl'}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Error</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Alert status="error">
              <AlertIcon />
              <AlertTitle>Error: </AlertTitle>
              <AlertDescription>
                {msgAlert}
              </AlertDescription>
            </Alert>
          </ModalBody>
        </ModalContent>
      </Modal>
    </>
  );
};

type LoginProps = {
  namePart: string;
  onClose: () => void;
  fetchExplorer?: () => Promise<void>;
  setIsDisabledLogout?: React.Dispatch<React.SetStateAction<boolean>>;
};

export default Login;
