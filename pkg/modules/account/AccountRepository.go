package account

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/laxeder/go-shop-service/pkg/modules/logger"
	"github.com/laxeder/go-shop-service/pkg/modules/redisdb"
	"github.com/laxeder/go-shop-service/pkg/modules/str"
)

var redisClient *redis.Client

func Repository() *Account {
	return &Account{}
}

func (a *Account) Save(account *Account) (err error) {
	var log = logger.New()

	ctx := context.Background()
	err = nil

	document := str.DocumentPad(account.Document)

	redisClient, err = redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return
	}

	key := fmt.Sprintf("accounts:%v", document)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "uuid", account.Uuid)
		rdb.HSet(ctx, key, "document", account.Document)
		rdb.HSet(ctx, key, "nickname", account.Nickname)
		rdb.HSet(ctx, key, "profession", account.Profession)
		rdb.HSet(ctx, key, "rg", account.RG)
		rdb.HSet(ctx, key, "birthday", account.Birthday)
		rdb.HSet(ctx, key, "gender", string(account.Gender))
		rdb.HSet(ctx, key, "created_at", account.CreatedAt)
		rdb.HSet(ctx, key, "updated_at", account.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível inserir um aconta no redis. (%v)", document)
		return
	}

	return
}

func (a *Account) Update(account *Account) (err error) {
	var log = logger.New()

	ctx := context.Background()
	err = nil

	document := str.DocumentPad(account.Document)

	redisClient, err = redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return
	}

	key := fmt.Sprintf("accounts:%v", document)

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
		log.Error().Err(err).Msgf("Não foi possível atualizar a conta no redis. (%v)", document)
		return
	}

	return
}

func (a *Account) GetDocument(document string) (*Account, error) {

	var log = logger.New()

	ctx := context.Background()

	a.Document = str.DocumentPad(document)

	redisClient, err := redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return a, err
	}

	key := fmt.Sprintf("accounts:%v", a.Document)
	res := redisClient.HMGet(ctx, key, "uuid", "document", "optins_cmr", "optins_cdu")
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o account com documento: %v.", a.Document)
		return a, err
	}

	err = res.Scan(a)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um account válido. (%v)", a.Document)
		return a, err
	}

	return a, err
}

func (a *Account) GetByDocument(document string) (*Account, error) {

	var log = logger.New()

	ctx := context.Background()

	document = str.DocumentPad(document)
	a.Document = document

	redisClient, err := redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return a, err
	}

	key := fmt.Sprintf("accounts:%v", document)
	res := redisClient.HGetAll(ctx, key)
	err = res.Err()
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível encontrar o account com documento: %v.", document)
		return a, err
	}

	err = res.Scan(a)
	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível mapear um account válido. (%v)", document)
		return a, err
	}

	return a, err
}

func (a *Account) SaveOptins(account *Account) (err error) {
	var log = logger.New()

	ctx := context.Background()
	err = nil

	document := str.DocumentPad(account.Document)

	redisClient, err = redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return
	}

	key := fmt.Sprintf("accounts:%v", document)

	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "optins_cdu", account.Options)
		rdb.HSet(ctx, key, "updated_at", account.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar o optins no redis. (%v)", document)
		return
	}

	return
}

func (u *Account) SaveDocument(dcm string, account *Account) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	document := str.DocumentPad(account.Document)

	redisClient, err = redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return err
	}

	key := fmt.Sprintf("accounts:%v", dcm)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "document", account.Document)
		rdb.HSet(ctx, key, "nickname", account.Nickname)
		rdb.HSet(ctx, key, "profession", account.Profession)
		rdb.HSet(ctx, key, "rg", account.RG)
		rdb.HSet(ctx, key, "birthday", account.Birthday)
		rdb.HSet(ctx, key, "gender", string(account.Gender))
		rdb.HSet(ctx, key, "updated_at", account.UpdatedAt)
		return err
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível atualizar o usuário com documento %v no redis.", document)
		return err
	}

	return nil
}

func (u *Account) Delete(account *Account) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	document := str.DocumentPad(account.Document)

	redisClient, err = redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return err
	}

	account.Status = Disabled

	key := fmt.Sprintf("accounts:%v", document)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(account.Status))
		rdb.HSet(ctx, key, "updated_at", account.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível deletar o conta com documento %v no redis.", document)
		return err
	}

	return nil
}

func (u *Account) Restore(account *Account) (err error) {

	var log = logger.New()

	err = nil

	ctx := context.Background()
	document := str.DocumentPad(account.Document)

	redisClient, err = redisdb.New(redisdb.AccountDatabase)
	if err != nil {
		log.Error().Err(err).Msgf("Erro ao acessar banco de dados (%v)", redisdb.AccountDatabase)
		return err
	}

	account.Status = Enabled

	key := fmt.Sprintf("accounts:%v", document)
	_, err = redisClient.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		rdb.HSet(ctx, key, "status", string(account.Status))
		rdb.HSet(ctx, key, "updated_at", account.UpdatedAt)
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msgf("Não foi possível restaurar a conta com documento %v no redis.", document)
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
		document := iter.Val()
		account, uErr := u.GetByDocument(document)
		if uErr != nil {
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
