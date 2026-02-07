//SPDX-License-Identifier:MIT

pragma solidity^0.8.10;

import "@V4-Core/src/interfaces/IHooks.sol";
import "@V4-Core/src/interfaces/IPoolManager.sol";

contract LPPrivacy is IHooks{
    IPoolManager public poolManager;

    uint256 public delay_block;

    event queuedIntent(uint256);
    event executedIntent(uint256);

    enum actionType {
            Add, Remove
        };

    bytes hookData;

    struct LiquidityIntent {
        address lp;
        PoolKey poolKey;
        actionType action;
        int24 tickLower; 
        int24 tickUpper;
        int128 liquidityDelta;
        uint256 queueBlock;
        uint256 executeAfterBlock;
        bool isExecuted;
    }

    mapping(uint256 => LiquidityIntent ) public intent;

    uint256 intentid = 0;

     function beforeAddLiquidity(
        address sender,
        PoolKey calldata key,
        ModifyLiquidityParams calldata params,
        bytes calldata hookData
    ) external returns (bytes4) {
        require( msg.sender == address(poolManager),"You can't proceed");
        LiquidityIntent memory _liquidityIntent ;
        PoolKey poolkey = key;
        _liquidityIntent.lp = sender;
        _liquidityIntent.poolKey = poolkey;
        _liquidityIntent.action = actionType.Add;
        _liquidityIntent.queueBlock = block.number;
        _liquidityIntent.executeAfterBlock = _liquidityIntent.queueBlock + delay_block;
        _liquidityIntent.isExecuted = false;
        intent[intentid] =  _liquidityIntent;
        intent[intentid].tickLower = params.tickLower;
        intent[intentid].tickUpper = params.tickUpper;
        intent[intentid].liquidityDelta = params.liquidityDelta;
        emit queuedIntent(intentid);
        intentid += 1;
        return this.beforeAddLiquidity.selector;
    }

     function beforeRemoveLiquidity(
        address sender,
        PoolKey calldata key,
        ModifyLiquidityParams calldata params,
        bytes calldata hookData
    ) external returns (bytes4) {
         LiquidityIntent memory _liquidityIntent ;
        PoolKey poolkey = key;
        _liquidityIntent.lp = sender;
        _liquidityIntent.poolKey = poolkey;
        _liquidityIntent.action = actionType.Remove;
        _liquidityIntent.queueBlock = block.number;
        _liquidityIntent.executeAfterBlock = _liquidityIntent.queueBlock + delay_block;
        _liquidityIntent.isExecuted = false;
        intent[intentid] =  _liquidityIntent;
        intent[intentid].tickLower = params.tickLower;
        intent[intentid].tickUpper = params.tickUpper;
        intent[intentid].liquidityDelta = params.liquidityDelta;
        emit queuedIntent(intentid);
        intentid += 1;
        return this.beforeRemoveLiquidity.selector;

    }

    function executeIntent(uint256 intentId) public {
        if (intentId >= intentid) {
            revert();
        }
        if (intent[intentId].isExecuted == true) {
            revert();
        }
        uint256 current_block = block.number;
        if(current_block < intent[intentId].executeAfterBlock) {
            revert();
        }
        ModifyLiquidityParams tParams;
        int24 tLower = intent[intentId].tickLower;
        int24 tUpper = intent[intentId].tickUpper;
        int128 lDelta = intent[intentId].liquidityDelta;
        tParams.tickLower = tLower;
        tParams.tickUpper = tUpper;
        tParams.liquidityDelta = lDelta;
        if (poolManager.modifyLiquidity(intent[intentId].poolKey, tParams, hookData)) {

        } else {
            revert();
        }
        if (current_block >= intent[intentId].executeAfterBlock) {
            intent[intentId].isExecuted = true;
        }

        emit executedIntent(intentId);
    }
}