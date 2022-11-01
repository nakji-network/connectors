use near_lake_framework::near_indexer_primitives::views::ActionView;

use crate::near_proto::{
    action::{self, AddKey, CreateAccount, DeleteAccount, DeleteKey, FunctionCall, Stake},
    AccessKey, Action,
};
use action::{DeployContract, Transfer};

pub mod block;
pub mod handle;
pub mod outcome;
pub mod receipt;
pub mod transaction;

fn action_from_view(action: ActionView) -> Action {
    match action {
        ActionView::CreateAccount => Action {
            action: Some(action::Action::CreateAccount(CreateAccount {})),
        },
        ActionView::DeployContract { code } => Action {
            action: Some(action::Action::DeployContract(DeployContract {
                code: code.to_string(),
            })),
        },
        ActionView::FunctionCall {
            method_name,
            args,
            gas,
            deposit,
        } => Action {
            action: Some(action::Action::FunctionCall(FunctionCall {
                method_name: method_name.to_string(),
                args: args.to_string(),
                gas,
                deposit: deposit.to_ne_bytes().to_vec(),
            })),
        },
        ActionView::Transfer { deposit } => Action {
            action: Some(action::Action::Transfer(Transfer {
                deposit: deposit.to_ne_bytes().to_vec(),
            })),
        },
        ActionView::Stake { stake, public_key } => Action {
            action: Some(action::Action::Stake(Stake {
                stake: stake.to_ne_bytes().to_vec(),
                public_key: public_key.to_string(),
            })),
        },
        ActionView::AddKey {
            public_key,
            access_key,
        } => Action {
            action: Some(action::Action::AddKey(AddKey {
                public_key: public_key.to_string(),
                access_key: Some(AccessKey::from(access_key)),
            })),
        },
        ActionView::DeleteKey { public_key } => Action {
            action: Some(action::Action::DeleteKey(DeleteKey {
                public_key: public_key.to_string(),
            })),
        },
        ActionView::DeleteAccount { beneficiary_id } => Action {
            action: Some(action::Action::DeleteAccount(DeleteAccount {
                beneficiary_id: beneficiary_id.to_string(),
            })),
        },
    }
}
