/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import {
  Contract,
  ContractFactory,
  ContractTransactionResponse,
  Interface,
} from "ethers";
import type {
  Signer,
  AddressLike,
  ContractDeployTransaction,
  ContractRunner,
} from "ethers";
import type { NonPayableOverrides } from "../../common";
import type {
  IbcMwUser,
  IbcMwUserInterface,
} from "../../IbcMiddleware.sol/IbcMwUser";

const _abi = [
  {
    type: "constructor",
    inputs: [
      {
        name: "_middleware",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "nonpayable",
  },
  {
    type: "receive",
    stateMutability: "payable",
  },
  {
    type: "function",
    name: "authorizeMiddleware",
    inputs: [
      {
        name: "middleware",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "authorizedMws",
    inputs: [
      {
        name: "",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [
      {
        name: "",
        type: "bool",
        internalType: "bool",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "mw",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "owner",
    inputs: [],
    outputs: [
      {
        name: "",
        type: "address",
        internalType: "address",
      },
    ],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "renounceOwnership",
    inputs: [],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "setDefaultMw",
    inputs: [
      {
        name: "_middleware",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "transferOwnership",
    inputs: [
      {
        name: "newOwner",
        type: "address",
        internalType: "address",
      },
    ],
    outputs: [],
    stateMutability: "nonpayable",
  },
  {
    type: "event",
    name: "OwnershipTransferred",
    inputs: [
      {
        name: "previousOwner",
        type: "address",
        indexed: true,
        internalType: "address",
      },
      {
        name: "newOwner",
        type: "address",
        indexed: true,
        internalType: "address",
      },
    ],
    anonymous: false,
  },
  {
    type: "error",
    name: "UnauthorizedIbcMiddleware",
    inputs: [],
  },
  {
    type: "error",
    name: "ackAddressMismatch",
    inputs: [],
  },
  {
    type: "error",
    name: "ackDataTooShort",
    inputs: [],
  },
] as const;

const _bytecode =
  "0x608060405234801561001057600080fd5b506040516104a73803806104a783398101604081905261002f916100d1565b61003833610081565b600180546001600160a01b0319166001600160a01b03831617905561007b816001600160a01b03166000908152600260205260409020805460ff19166001179055565b50610101565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100e357600080fd5b81516001600160a01b03811681146100fa57600080fd5b9392505050565b610397806101106000396000f3fe6080604052600436106100735760003560e01c8063715018a61161004e578063715018a6146101065780638da5cb5b1461011b578063a742d78c1461014d578063f2fde38b1461016d57600080fd5b8062e82cef1461007f5780633a7fbcbd146100a15780633b90b042146100e657600080fd5b3661007a57005b600080fd5b34801561008b57600080fd5b5061009f61009a366004610331565b61018d565b005b3480156100ad57600080fd5b506100d16100bc366004610331565b60026020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b3480156100f257600080fd5b5061009f610101366004610331565b6101c0565b34801561011257600080fd5b5061009f6101d4565b34801561012757600080fd5b506000546001600160a01b03165b6040516001600160a01b0390911681526020016100dd565b34801561015957600080fd5b50600154610135906001600160a01b031681565b34801561017957600080fd5b5061009f610188366004610331565b6101e8565b610195610263565b61019e816102bd565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6101c8610263565b6101d1816102bd565b50565b6101dc610263565b6101e660006102e1565b565b6101f0610263565b6001600160a01b03811661025a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b6101d1816102e1565b6000546001600160a01b031633146101e65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610251565b6001600160a01b03166000908152600260205260409020805460ff19166001179055565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006020828403121561034357600080fd5b81356001600160a01b038116811461035a57600080fd5b939250505056fea26469706673582212207aba41329608ee1c3c9569c89f75fa777781ac981a73f0254c963373fa2caa3664736f6c634300080f0033";

type IbcMwUserConstructorParams =
  | [signer?: Signer]
  | ConstructorParameters<typeof ContractFactory>;

const isSuperArgs = (
  xs: IbcMwUserConstructorParams
): xs is ConstructorParameters<typeof ContractFactory> => xs.length > 1;

export class IbcMwUser__factory extends ContractFactory {
  constructor(...args: IbcMwUserConstructorParams) {
    if (isSuperArgs(args)) {
      super(...args);
    } else {
      super(_abi, _bytecode, args[0]);
    }
  }

  override getDeployTransaction(
    _middleware: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ): Promise<ContractDeployTransaction> {
    return super.getDeployTransaction(_middleware, overrides || {});
  }
  override deploy(
    _middleware: AddressLike,
    overrides?: NonPayableOverrides & { from?: string }
  ) {
    return super.deploy(_middleware, overrides || {}) as Promise<
      IbcMwUser & {
        deploymentTransaction(): ContractTransactionResponse;
      }
    >;
  }
  override connect(runner: ContractRunner | null): IbcMwUser__factory {
    return super.connect(runner) as IbcMwUser__factory;
  }

  static readonly bytecode = _bytecode;
  static readonly abi = _abi;
  static createInterface(): IbcMwUserInterface {
    return new Interface(_abi) as IbcMwUserInterface;
  }
  static connect(address: string, runner?: ContractRunner | null): IbcMwUser {
    return new Contract(address, _abi, runner) as unknown as IbcMwUser;
  }
}