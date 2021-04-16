use seata_order;
delete from undo_log;
delete from branch_transaction;
delete from so_a;

use seata_product;
delete from undo_log;
delete from branch_transaction;
delete from so_b;

use seata;
delete from branch_table;
delete from global_table;
delete from lock_table;
