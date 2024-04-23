FLARE_SMART_CONTRACTS_V2 = ../flare-smart-contracts-v2
GO_CLIENT_RELATIVE = ../go-client
SOLC_OPTIONS = @openzeppelin/=node_modules/@openzeppelin/ flare-smart-contracts/=node_modules/flare-smart-contracts/ --evm-version london --optimize --optimize-runs 200
CONTRACTS_FOLDER = contracts-interface

compile: compile-solc compile-abigen 

compile-solc: compile-solc-base compile-solc-mock
compile-solc-base: compile-solc-fast_updater-abi compile-solc-fast_updater-bin compile-solc-fast_updates-configuration-abi compile-solc-fast_updates-configuration-bin compile-solc-fast_update-incentive-abi compile-solc-fast_update-incentive-bin compile-solc-submission-abi
compile-solc-mock: compile-solc-mock-abi compile-solc-mock-bin

compile-solc-fast_updater-abi:
	cd $(FLARE_SMART_CONTRACTS_V2); solc $(SOLC_OPTIONS) --abi contracts/fastUpdates/implementation/FastUpdater.sol -o $(GO_CLIENT_RELATIVE)/$(CONTRACTS_FOLDER)  --overwrite

compile-solc-fast_updater-bin:
	cd $(FLARE_SMART_CONTRACTS_V2); solc $(SOLC_OPTIONS) --bin contracts/fastUpdates/implementation/FastUpdater.sol -o $(GO_CLIENT_RELATIVE)/$(CONTRACTS_FOLDER)  --overwrite

compile-solc-fast_updates-configuration-abi:
	cd $(FLARE_SMART_CONTRACTS_V2); solc $(SOLC_OPTIONS) --abi contracts/fastUpdates/implementation/FastUpdatesConfiguration.sol -o $(GO_CLIENT_RELATIVE)/$(CONTRACTS_FOLDER)  --overwrite

compile-solc-fast_updates-configuration-bin:
	cd $(FLARE_SMART_CONTRACTS_V2); solc $(SOLC_OPTIONS) --bin contracts/fastUpdates/implementation/FastUpdatesConfiguration.sol -o $(GO_CLIENT_RELATIVE)/$(CONTRACTS_FOLDER)  --overwrite

compile-solc-fast_update-incentive-abi:
	cd $(FLARE_SMART_CONTRACTS_V2); solc $(SOLC_OPTIONS) --abi contracts/fastUpdates/implementation/FastUpdateIncentiveManager.sol -o $(GO_CLIENT_RELATIVE)/$(CONTRACTS_FOLDER)  --overwrite

compile-solc-fast_update-incentive-bin:
	cd $(FLARE_SMART_CONTRACTS_V2); solc $(SOLC_OPTIONS) --bin contracts/fastUpdates/implementation/FastUpdateIncentiveManager.sol -o $(GO_CLIENT_RELATIVE)/$(CONTRACTS_FOLDER)  --overwrite

compile-solc-submission-abi:
	cd $(FLARE_SMART_CONTRACTS_V2); solc $(SOLC_OPTIONS) --abi contracts/protocol/implementation/Submission.sol -o $(GO_CLIENT_RELATIVE)/$(CONTRACTS_FOLDER)  --overwrite

compile-solc-mock-abi:
	cd $(FLARE_SMART_CONTRACTS_V2); solc $(SOLC_OPTIONS) --abi contracts/fastUpdates/mock/FlareSystemMock.sol -o $(GO_CLIENT_RELATIVE)/$(CONTRACTS_FOLDER)  --overwrite

compile-solc-mock-bin:
	cd $(FLARE_SMART_CONTRACTS_V2); solc $(SOLC_OPTIONS) --bin contracts/fastUpdates/mock/FlareSystemMock.sol -o $(GO_CLIENT_RELATIVE)/$(CONTRACTS_FOLDER)  --overwrite


compile-abigen: compile-abigen-flare_system_manager compile-abigen-flare_mock compile-abigen-voter_registry compile-abigen-fast_updater compile-abigen-fast_update-configuration compile-abigen-fast_update-incentive compile-abigen-submission

compile-abigen-flare_system_manager:
	cd $(CONTRACTS_FOLDER); abigen --abi=IFlareSystemsManager.abi --pkg=system_manager --out=system_manager/FlareSystemManager.go

compile-abigen-flare_mock:
	cd $(CONTRACTS_FOLDER); abigen --bin=FlareSystemMock.bin --abi=FlareSystemMock.abi --pkg=mock --out=mock/FlareSystemMock.go

compile-abigen-voter_registry:
	cd $(CONTRACTS_FOLDER); abigen --abi=IVoterRegistry.abi --pkg=voter_registry --out=voter_registry/VoterRegistry.go

compile-abigen-fast_updater:
	cd $(CONTRACTS_FOLDER); abigen --bin=FastUpdater.bin --abi=FastUpdater.abi --pkg=fast_updater --out=fast_updater/FastUpdater.go

compile-abigen-fast_update-configuration:
	cd $(CONTRACTS_FOLDER); abigen --bin=FastUpdatesConfiguration.bin --abi=FastUpdatesConfiguration.abi --pkg=fast_updates_configuration --out=fast_updates_configuration/FastUpdatesConfiguration.go

compile-abigen-fast_update-incentive:
	cd $(CONTRACTS_FOLDER); abigen --bin=FastUpdateIncentiveManager.bin --abi=FastUpdateIncentiveManager.abi --pkg=incentive --out=incentive/FastUpdateIncentiveManager.go

compile-abigen-submission:
	cd $(CONTRACTS_FOLDER); abigen --abi=Submission.abi --pkg=submission --out=submission/Submission.go
