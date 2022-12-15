package account

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/date"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
	"github.com/laxeder/go-shop-service/pkg/modules/user"
)

var redisClient *redis.Client

func Repository() *Account {
	return &Account{}
}

func (a *Account) Save(account *Account) (err error) {
	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err = redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return
	}

	key := fmt.Sprintf("accounts:%v", account.Uid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "uuid", account.Uuid)
		rdb.HSet(ctx, key, "uid", account.Uid)
		rdb.HSet(ctx, key, "nickname", account.Nickname)
		rdb.HSet(ctx, key, "profession", account.Profession)
		rdb.HSet(ctx, key, "rg", account.RG)
		rdb.HSet(ctx, key, "birthday", account.Birthday)
		rdb.HSet(ctx, key, "gender", string(account.Gender))
		rdb.HSet(ctx, key, "status", string(account.Status))
		rdb.HSet(ctx, key, "created_at", account.CreatedAt)
		rdb.HSet(ctx, key, "updated_at", account.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível inserir uma conta com o uid no redis. (%v)", account.Uid)
		return
	}

	// carrega o usuário da base de dados para atualizar as contas
	userDatabase, err := user.Repository().GetByUuid(account.Uuid)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar validar usuário (%v), (%v).", account.Uid, err)
		return
	}

	userDatabase.Accounts = append(userDatabase.Accounts, account.Uid)
	userDatabase.UpdatedAt = date.NowUTC()

	err = user.Repository().Update(userDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao tentar atualizar contas do usuário (%v), (%v).", account.Uid, err)
		return
	}

	return
}

func (a *Account) Update(account *Account) (err error) {
	var log = logger.New()

	ctx := context.Background()
	err = nil

	redisClient, err = redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return
	}

	key := fmt.Sprintf("accounts:%v", account.Uid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "nickname", account.Nickname)
		rdb.HSet(ctx, key, "profession", account.Profession)
		rdb.HSet(ctx, key, "rg", account.RG)
		rdb.HSet(ctx, key, "birthday", account.Birthday)
		rdb.HSet(ctx, key, "gender", string(account.Gender))
		rdb.HSet(ctx, key, "updated_at", account.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar a conta no redis. (%v)", account.Uid)
		return
	}

	return
}

func (u *Account) GetUid(uid string) (account *Account, err error) {

	var log = logger.New()

	account = &Account{}
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return nil, err
	}

	key := fmt.Sprintf("accounts:%v", uid)
	res := redisClient.HMGet(ctx, key, "uuid", "uid", "status")
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar a conta com uid: %v.", uid)
		return u, err
	}

	err = res.Scan(account)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear uma conta válida para o uid %v.", uid)
		return nil, err
	}

	if account.Status == Disabled {
		account = &Account{Status: Disabled}
		return account, nil
	}

	return account, nil
}

func (a *Account) GetByUid(uid string) (account *Account, err error) {

	var log = logger.New()

	account = &Account{}
	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return nil, err
	}

	key := fmt.Sprintf("accounts:%v", uid)
	res := redisClient.HGetAll(ctx, key)

	err = res.Err()

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar a conta com uid: %v.", uid)
		return nil, err
	}

	err = res.Scan(account)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear uma conta válida para o uid %v.", uid)
		return nil, err
	}

	if account.Status == Disabled {
		account = &Account{Status: Disabled}
		return account, nil
	}

	return account, nil
}

func (a *Account) SaveOptins(account *Account) (err error) {
	var log = logger.New()

	ctx := context.Background()
	err = nil

	redisClient, err = redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return
	}

	key := fmt.Sprintf("accounts:%v", account.Uid)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "optins_cdu", account.Options)
		rdb.HSet(ctx, key, "updated_at", account.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar o optins no redis. (%v)", account.Uid)
		return
	}

	return
}

func (u *Account) Delete(account *Account) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err = redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return err
	}

	account.Status = Disabled

	key := fmt.Sprintf("accounts:%v", account.Uid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(account.Status))
		rdb.HSet(ctx, key, "updated_at", account.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível deletar o conta com uid %v no redis.", account.Uid)
		return err
	}

	return nil
}

func (u *Account) Restore(account *Account) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err = redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return err
	}

	account.Status = Enabled

	key := fmt.Sprintf("accounts:%v", account.Uid)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(account.Status))
		rdb.HSet(ctx, key, "updated_at", account.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível restaurar a conta com uid %v no redis.", account.Uid)
		return err
	}

	return nil
}

func (u *Account) GetList() (accounts []Account, err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()

	redisClient, err := redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return nil, err
	}

	iter := redisClient.Scan(ctx, 0, "accounts:*", 0).Iterator()
	for iter.Next(ctx) {
		uid := strings.Replace(iter.Val(), "accounts:", "", 2)
		account, aErr := u.GetByUid(uid)

		if aErr != nil {
			continue
		}

		if account.Status == Disabled {
			continue
		}

		accounts = append(accounts, *account)
	}

	err = iter.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível listar as contas do banco de dados.")
		return nil, err
	}

	return accounts, nil
}
