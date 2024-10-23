#!/usr/bin/env node
import { ContractRegistryLoader, parseObjFromFile } from "..";
import { loadContractUpdateRegistry } from "../evm/schemas/contractUpdate";
import { updateContractsForChain } from "../updateContract";

import { getOutputLogger } from "../utils/cli";
import { UPDATE_SPECS_PATH } from "../utils/constants";
import { parseArgsFromCLI } from "../utils/io";

async function main() {
  const { chain, accounts, args, extraBindingsPath, externalContractsPath } =
    await parseArgsFromCLI();
  const updateSpecs = (args.UPDATE_SPECS_PATH as string) || UPDATE_SPECS_PATH;

  const contractUpdates = loadContractUpdateRegistry(
    parseObjFromFile(updateSpecs)
  );

  let extraContractFactories: Record<string, object> | null = null;
  if (extraBindingsPath) {
    try {
      extraContractFactories = await import(extraBindingsPath);
    } catch (e) {
      throw new Error(
        `Failed to import extra contract factories from ${extraBindingsPath}: ${e}`
      );
    }
  }

  const existingContracts = externalContractsPath
    ? ContractRegistryLoader.loadSingle(externalContractsPath)
    : ContractRegistryLoader.emptySingle();

  updateContractsForChain(
    chain,
    accounts.mustGet(chain.chainName),
    existingContracts,
    contractUpdates,
    getOutputLogger(),
    {
      dryRun: false,
      forceDeployNewContracts: false,
      writeContracts: true,
      extraContractFactories: extraContractFactories ?? undefined,
    }
  );
}
main();
